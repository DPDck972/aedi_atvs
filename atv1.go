package main

import (
	"errors"
	"fmt"
)

//Struct de tipo abstrato

type list interface {
	size() int
	get(index int) (int, error)
	push_back(e float32)
	addOnIndex(index int, e float32) error
	remove(index int) error
	pop() error
	set(e float32, index int)
}

type ArrayList struct {
	array     []float32
	elementos int
}

type Noh struct {
	valor    float32
	next     *Noh
	previous *Noh
}

type DoubleLinkedList struct {
	head      *Noh
	tail      *Noh
	elementos int
}

//Metodos

// Iniciar Arraylist (Como usa array)
func (l *ArrayList) Init(s int) {
	l.array = make([]float32, s)
}

// Função get size
func (lista *ArrayList) Size() int {
	return lista.elementos
}

func (lista *DoubleLinkedList) Size() int {
	return lista.elementos
}

// Função get element
func (lista *ArrayList) get(index int) (float32, error) {
	if index < 0 || index >= lista.elementos {
		return -1, errors.New(fmt.Sprintf("Index invalido: %d", index))
	} else {
		return lista.array[index], nil
	}
}

func (lista *DoubleLinkedList) get(index int) (float32, error) {
	if index < 0 || index >= lista.elementos {
		return -1, errors.New(fmt.Sprintf("Index invalido: %d", index))
	}
	if lista.head == nil || lista.tail == nil {
		return -1, errors.New(fmt.Sprintf("Lista vazia"))
	}
	var aux *Noh
	if index < (lista.elementos / 2) {
		aux = lista.head
		for x := 0; x < index; x++ {
			if aux.next == nil {
				return -1, fmt.Errorf("Erro ao acessar a lista")
			}
			aux = aux.next
		}
	} else {
		aux = lista.tail
		for x := lista.elementos - 1; x > index; x-- {
			if aux.previous == nil {
				return -1, fmt.Errorf("Erro ao acessar a lista")
			}
			aux = aux.previous
		}
	}
	return aux.valor, nil
}

// Função extender tamanho do vetor
func (lista *ArrayList) extend(tam int) {
	new_array := make([]float32, tam*lista.elementos)
	for i := 0; i < lista.elementos; i++ {
		new_array[i] = lista.array[i]
	}
	lista.array = new_array
}

// Funções de inserir no final
func (lista *ArrayList) push_back(e float32) {
	if lista.elementos == len(lista.array) {
		lista.extend(2)
	}
	lista.array[lista.elementos] = e
	lista.elementos++
}

func (lista *DoubleLinkedList) push_back(e float32) {
	novo_noh := &Noh{valor: e, next: nil, previous: lista.head}
	if lista.elementos == 0 {
		lista.head = novo_noh
		lista.tail = novo_noh
	} else {
		novo_noh.previous = lista.tail
		lista.tail.next = novo_noh
		lista.tail = novo_noh
	}
	lista.elementos++
}

// Função inserir no inicio
func (lista *DoubleLinkedList) push_front(e float32) {
	novo_noh := &Noh{valor: e, next: lista.head, previous: nil}
	lista.head = novo_noh
	lista.elementos++
}

// Função adicionar em um index especifico
func (lista *ArrayList) addOnIndex(index int, e float32) error {
	if index < 0 || index >= lista.elementos {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
	if lista.elementos == len(lista.array) {
		lista.extend(2)
	}
	for i := lista.elementos; i > index; i-- {
		lista.array[i] = lista.array[i-1]
	}
	lista.array[index] = e
	lista.elementos++
	return nil
}

func (lista *DoubleLinkedList) addOnIndex(index int, e float32) error {
	if index < 0 || index >= lista.elementos {
		return errors.New(fmt.Sprintf("Index inválido: %d", index))
	}
	novo_noh := &Noh{valor: e, next: nil, previous: nil}
	var aux *Noh
	if index == 0 {
		novo_noh.next = lista.head
		lista.head.previous = novo_noh

	} else if index < (lista.elementos / 2) {
		aux = lista.head
		for i := 0; i <= index; i++ {
			if aux.next == nil {
				return fmt.Errorf("Erro ao acessar a lista")
			}
			aux = aux.next
		}
		novo_noh.next = aux.next
		novo_noh.previous = aux
		aux.next = novo_noh
	} else {
		aux = lista.tail
		for i := lista.elementos; i >= index; i++ {
			if aux.previous == nil {
				return fmt.Errorf("Erro ao acessar a lista")
			}
			aux = aux.previous
		}
		novo_noh.previous = aux.previous
		novo_noh.next = aux
		aux.previous = novo_noh
	}
	lista.elementos++
	return nil
}

func main() {
	dll := &DoubleLinkedList{}
	for i := 1; i <= 50; i++ {
		dll.push_back(float32(i))
	}
	val, erro := dll.get(0)
	if erro != nil {
		fmt.Println("Erro:", erro)
	}
	fmt.Println(val)
}
