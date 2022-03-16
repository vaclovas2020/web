/* Copyright (c) 2022 Vaclovas Lapinskis. All rights reserved */

package web

/* Yaml main  struct */
type Config struct {
	Version string       `yaml:"version"`
	Project Project      `yaml:"project"`
	Config  []ConfigItem `yaml:"config"`
}

/* Yaml project struct */
type Project struct {
	Name        string      `yaml:"name"`
	Version     string      `yaml:"version"`
	Directories Directories `yaml:"directories"`
}

/* Yaml directories struct */
type Directories struct {
	SourceDir   string `yaml:"sourceDir"`
	BytecodeDir string `yaml:"bytecodeDir"`
}

/* Yaml config array item */
type ConfigItem struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}
