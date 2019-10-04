package baddependency

import (
       "fmt"
       "github.com/jfrog-solutiontest/modtest"
)

func ModuleTestPackageName () {
   fmt.Println ("Hello World")
   modtest.ModuleTestPackageName()
}
