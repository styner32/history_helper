package history_helper_test

import (
	. "github.com/styner32/history_helper"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Trie", func() {
	Describe("CreateTrie", func() {
		It("creates a node", func() {
			node := CreateTrie()
			Expect(node.Value).To(Equal(rune(0)))
			Expect(node.Depth).To(Equal(0))
			Expect(node.Parent).Should(BeNil())
			Expect(node.Next).Should(BeNil())
			Expect(node.Child).Should(BeNil())
		})
	})

	Describe("AddWord", func() {
		It("adds nodes to trie", func() {
			node := CreateTrie()
			node.AddWord("vi")
			child := node.Child
			Expect(child.Value).To(Equal('v'))
			Expect(child.Depth).To(Equal(1))
			Expect(child.Child.Value).To(Equal('i'))
			Expect(child.Child.Depth).To(Equal(2))
			Expect(child.Parent).To(Equal(node))
		})

		It("shared parents that has same prefix", func() {
			node := CreateTrie()
			node.AddWord("vm")
			node.AddWord("vx")
			child := node.Child
			Expect(child.Value).To(Equal('v'))
			Expect(child.Child.Value).To(Equal('m'))
			Expect(child.Child.Parent).To(Equal(child))
			Expect(child.Child.Next.Value).To(Equal('x'))
			Expect(child.Child.Next.Parent).To(Equal(child))
		})
	})
})
