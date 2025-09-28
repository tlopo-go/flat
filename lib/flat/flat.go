package flat

import (
	"fmt"
	"github.com/iancoleman/orderedmap"
	"gopkg.in/yaml.v3"
	"strings"
)

func load(input string) (result yaml.Node, err error) {
	err = yaml.Unmarshal([]byte(input), &result)
	return
}

func keys(node *yaml.Node) (keys *[]interface{}) {
	keys = &[]interface{}{}
	length := len(node.Content)
	i := 0
	for i < length {
		*keys = append(*keys, node.Content[i].Value)
		i += 2
	}
	return
}

func values(node *yaml.Node) (values *[]interface{}) {
	values = &[]interface{}{}
	length := len(node.Content)
	i := 1
	for i < length {
		*values = append(*values, node.Content[i])
		i += 2
	}
	return
}

func formatKeys(list *[]interface{}, separator string) (result string) {
	r := make([]string, len(*list))
	for i, v := range *list {
		r[i] = fmt.Sprintf("%v", v)
	}
	result = strings.Join(r, separator)
	return
}

func pop(stack *[]interface{}) {
	//fmt.Printf("popping, stack: %#v\n", stack)
	if len(*stack) > 0 {
		*stack = (*stack)[:len(*stack)-1]
	}
}

func flatRec(node *yaml.Node, name string, stack *[]interface{}, result *orderedmap.OrderedMap, separator string) {
	if node.Kind == yaml.MappingNode && len(*(keys(node))) > 0 {
		// Get Keys
		nodeKeys := keys(node)
		nodeValues := values(node)
		// for each key, append key in stack
		for i := 0; i < len(*nodeKeys); i++ {
			// recurse with value, key name, stack
			k := (*nodeKeys)[i]
			v := (*nodeValues)[i]
			*stack = append(*stack, k)
			//fmt.Printf("stack: %#v, element: %#v counter:%#v len:%#v\n", stack, v, i, len(*nodeKeys))
			flatRec(v.(*yaml.Node), k.(string), stack, result, separator)
		}
		pop(stack)
	} else if node.Kind == yaml.SequenceNode && len(node.Content) > 0 {
		for i, e := range node.Content {
			// append index to stack
			*stack = append(*stack, i)
			// recurse with element, index,  stack
			flatRec(e, fmt.Sprintf("%d", i), stack, result, separator)
			//fmt.Printf("stack: %#v, element: %#v counter:%#v len:%#v\n", stack, e, i, len(node.Content))
		}
		pop(stack)
	} else /*if node.Kind == yaml.ScalarNode*/ {
		value := node.Value
		if node.Kind == yaml.MappingNode {
			value = "{}"
		} else if node.Kind == yaml.SequenceNode {
			value = "[]"
		}
		formatted_keys := formatKeys(stack, separator)
		(*result).Set(formatted_keys, value)
		//fmt.Printf("stack: %#v, element: %#v\n", stack, node.Value)
		pop(stack)
	}
}

func rec(node *yaml.Node, previous *yaml.Node, v *[]interface{}) {
	for _, n := range node.Content {
		rec(n, node, v)
	}

}

type Flat struct {
	content   string
	separator string
}

func New() *Flat {
	return &Flat{separator: " | "}
}

func (f *Flat) Content(content string) *Flat {
	f.content = content
	return f
}

func (f *Flat) Separator(separator string) *Flat {
	f.separator = separator
	return f
}

func (f *Flat) Run() (result *orderedmap.OrderedMap) {
	node, _ := load(string(f.content))
	rec(&node, nil, &[]interface{}{})
	result = orderedmap.New()
	flatRec(node.Content[0], "", &[]interface{}{}, result, f.separator)
	return
}
