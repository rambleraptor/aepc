// Copyright 2023 Yusuke Fredrick Tsutsumi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"

	"github.com/aep-dev/aepc/loader"
	"github.com/aep-dev/aepc/parser"
	"github.com/aep-dev/aepc/schema"
	"github.com/aep-dev/aepc/validator"
	"github.com/aep-dev/aepc/writer/proto"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewCommand() *cobra.Command {
	var inputFile string
	var outputFile string

	c := &cobra.Command{
		Use:   "aepc",
		Short: "aepc compiles resource representations to full proto rpcs",
		Long:  "aepc compiles resource representations to full proto rpcs",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: error handling
			s := &schema.Service{}
			input, err := readFile(inputFile)
			fmt.Printf("input: %s\n", string(input))
			if err != nil {
				log.Fatalf("unable to read file: %v", err)
			}
			ext := filepath.Ext(inputFile)
			err = unmarshal(ext, input, s)
			if err != nil {
				log.Fatal(err)
			}
			errors := validator.ValidateService(s)
			if len(errors) > 0 {
				log.Fatalf("error validating service: %v", errors)
			}
			ps, err := parser.NewParsedService(s)
			if err != nil {
				log.Fatal(err)
			}
			proto, _ := proto.WriteServiceToProto(ps)

			err = writeFile(outputFile, proto)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("output file: %s\n", outputFile)
			fmt.Printf("output proto: %s\n", proto)
		},
	}
	c.Flags().StringVarP(&inputFile, "input", "i", "", "input files with resource")
	c.Flags().StringVarP(&outputFile, "output", "o", "", "output file to use")
	return c
}

func unmarshal(ext string, b []byte, s *schema.Service) error {
	switch ext {
	case ".proto":
		if err := loader.ReadServiceFromProto(b, s); err != nil {
			return fmt.Errorf("unable to decode proto %q: %w", string(b), err)
		}
	case ".yaml":
		asJson, err := yaml.YAMLToJSON(b)
		if err != nil {
			return fmt.Errorf("unable to decode yaml to JSON %q: %w", string(b), err)
		}
		if err := protojson.Unmarshal(asJson, s); err != nil {
			log.Fatal(fmt.Errorf("unable to decode proto %q: %w", string(b), err))
		}
	case ".json":
		if err := protojson.Unmarshal(b, s); err != nil {
			return fmt.Errorf("unable to decode json %q: %w", string(b), err)
		}
	default:
		return fmt.Errorf("extension %v is unsupported", ext)
	}
	return nil
}

func readFile(fileName string) ([]byte, error) {
	var value []byte
	f, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	bytesRead := 1
	for bytesRead > 0 {
		readBytes := make([]byte, 10000)
		bytesRead, err = f.Read(readBytes)
		if bytesRead > 0 {
			value = append(value, readBytes[:bytesRead]...)
		}
		if err != io.EOF && err != nil {
			return nil, err
		}
	}
	err = f.Close()
	if err != nil {
		return nil, err
	}
	return value, nil
}

func writeFile(fileName string, value []byte) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(value)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}
	return nil
}
