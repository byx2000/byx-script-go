package interpreter

import (
	. "byx-script-go/common"
	. "byx-script-go/parser"
	"fmt"
	"path/filepath"
)

type RunConfig struct {
	Builtins    map[string]Value
	ImportPaths []string
}

// RunScript 执行ByxScript脚本
func RunScript(script string, config RunConfig) {
	scope := NewEmptyScope()

	// 添加默认内建函数
	addBuiltins(scope)

	// 添加自定义内建函数
	if config.Builtins != nil {
		for k, v := range config.Builtins {
			scope.DeclareVar(k, v)
		}
	}

	// 添加默认导入路径
	importPaths := config.ImportPaths
	if importPaths == nil {
		importPaths = []string{"./"}
	}

	// 解析脚本
	program := ParseScript(script)

	// 解析导入
	imports := parseImports(program.Imports, importPaths)

	// 计算加载顺序
	order := getLoadOrder(imports)

	// 按顺序执行依赖
	for _, name := range order {
		execute(imports[name], scope)
	}

	// 执行脚本
	execute(program, scope)
}

func parseImportName(importName string, importPaths []string) Program {
	for _, path := range importPaths {
		script, err := ReadFileAsString(filepath.Join(path, importName+".bs"))
		if err != nil {
			continue
		}
		return ParseScript(script)
	}
	panic(fmt.Sprintf("cannot resolve import name: %s", importName))
}

func parseImports(imports []string, importPaths []string) map[string]Program {
	result := map[string]Program{}
	namesToParse := []string{}
	namesToParse = append(namesToParse, imports...)
	for {
		cnt := len(namesToParse)
		if cnt == 0 {
			break
		}

		for i := 0; i < cnt; i++ {
			name := namesToParse[0]
			namesToParse = namesToParse[1:]
			p := parseImportName(name, importPaths)
			result[name] = p
			for _, n := range p.Imports {
				if _, exist := result[n]; !exist {
					namesToParse = append(namesToParse, n)
				}
			}
		}
	}

	return result
}

func getLoadOrder(imports map[string]Program) []string {
	// 计算依赖关系
	dependOn := map[string]map[string]bool{}
	for k, v := range imports {
		dependOn[k] = map[string]bool{}
		for _, n := range v.Imports {
			dependOn[k][n] = true
		}
	}

	// 计算反向依赖关系
	dependBy := map[string]map[string]bool{}
	for k, v := range dependOn {
		for n, _ := range v {
			set, exist := dependBy[n]
			if !exist {
				set = map[string]bool{}
			}
			set[k] = true
			dependBy[n] = set
		}
	}

	// 计算出度
	out := map[string]int{}
	for k, v := range dependOn {
		out[k] = len(v)
	}

	// 拓扑排序
	ready := map[string]bool{}
	for k, v := range out {
		if v == 0 {
			ready[k] = true
		}
	}

	order := []string{}
	for {
		if len(ready) == 0 {
			break
		}

		var n string
		for k, _ := range ready {
			n = k
			break
		}
		delete(ready, n)
		order = append(order, n)
		if ns, exist := dependBy[n]; exist {
			for n2, _ := range ns {
				o := out[n2]
				out[n2] = o - 1
				if o == 1 {
					ready[n2] = true
				}
			}
		}
	}

	// 检测循环依赖
	if len(order) != len(imports) {
		panic("circular dependency import detected")
	}

	return order
}

func addBuiltins(scope Scope) {
	scope.DeclareVar("print", Print)
	scope.DeclareVar("println", Println)

	scope.DeclareVar("isInteger", IsInteger)
	scope.DeclareVar("isDouble", IsDouble)
	scope.DeclareVar("isBool", IsBool)
	scope.DeclareVar("isString", IsString)
	scope.DeclareVar("isList", IsList)
	scope.DeclareVar("isCallable", IsCallable)
	scope.DeclareVar("isObject", IsObject)
	scope.DeclareVar("isUndefined", IsUndefined)
	scope.DeclareVar("hashcode", Hashcode)

	scope.DeclareVar("pow", Pow)
	scope.DeclareVar("sin", Sin)
	scope.DeclareVar("cos", Cos)
	scope.DeclareVar("tan", Tan)
	scope.DeclareVar("exp", Exp)
	scope.DeclareVar("ln", Ln)
	scope.DeclareVar("log10", Log10)
	scope.DeclareVar("sqrt", Sqrt)
	scope.DeclareVar("round", Round)
	scope.DeclareVar("ceil", Ceil)
	scope.DeclareVar("floor", Floor)
}
