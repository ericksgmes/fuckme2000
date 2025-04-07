package vm

import (
    "fmt"
    "strconv"
)

type VM struct {
    Stack []int
    Vars  map[string]int
}

func New() *VM {
    return &VM{
        Stack: []int{},
        Vars:  make(map[string]int),
    }
}

func (vm *VM) Run(program []Instruction) {
    for _, instr := range program {
        switch instr.Op {
        case "PUSH":
            val, err := strconv.Atoi(instr.Arg)
            if err != nil {
                panic("PUSH needs integer arg")
            }
            vm.Stack = append(vm.Stack, val)

        case "STORE":
            if len(vm.Stack) == 0 {
                panic("STORE with empty stack")
            }
            top := vm.pop()
            vm.Vars[instr.Arg] = top

        case "LOAD":
            val, ok := vm.Vars[instr.Arg]
            if !ok {
                panic("LOAD on undefined variable: " + instr.Arg)
            }
            vm.Stack = append(vm.Stack, val)

        case "ADD":
            if len(vm.Stack) < 2 {
                panic("ADD needs 2 values")
            }
            b := vm.pop()
            a := vm.pop()
            vm.Stack = append(vm.Stack, a+b)

        case "PRINT":
            val := vm.pop()
            printlnVal(fmt.Sprintf("%d", val))        

        default:
            panic("Unknown instruction: " + instr.Op)
        }
    }
}

func (vm *VM) pop() int {
    if len(vm.Stack) == 0 {
        panic("pop from empty stack")
    }
    val := vm.Stack[len(vm.Stack)-1]
    vm.Stack = vm.Stack[:len(vm.Stack)-1]
    return val
}


