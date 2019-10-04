package baddependcy

import (
       "fmt"
       "github.com/jfrog-solutiontest/modtest"
)

func baddependency() {
   fmt.Println ("Hello World")
   modtest.ModuleTestPackageName()
}
