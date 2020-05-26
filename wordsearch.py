"""
211. Add and Search Word - Data structure design

Design a data structure that supports the following two operations:
> void addWord(word)
> bool search(word)
search(word) can search a literal word or a regular expression string containing only letters a-z or .. 
A . means it can represent any one letter.

https://leetcode.com/problems/add-and-search-word-data-structure-design/
"""

class TrieNode:
    
    def __init__(self, char: str):
        self.char = char
        self.children = []
        self.full_word = False

class WordDictionary:

    def __init__(self):
        self.root = TrieNode("*")

    # addWord adds a word to the TrieNode tree
    def addWord(self, word: str) -> None:
        node = self.root
        
        for char in word:
            found = False
            
            for child in node.children:
                if child.char == char:
                    found = True
                    node = child
                    break
            
            if not found:
                new_node = TrieNode(char)
                node.children.append(new_node)
                node = new_node
        
        node.full_word = True
                
    
    # search finds the provided word 
    def search(self, word: str) -> bool:
        # deepSearch recursively searches for the provided word
        def deepSearch(word: str, index: int, node: TrieNode) -> bool:
            if node.full_word and index == len(word):
                return True
            elif index == len(word):
                return False

            for child in node.children:
                if (child.char == word[index] or word[index] == ".") and deepSearch(word, index + 1, child):
                    return True

            return False
        
        return deepSearch(word, 0, self.root)
    

    
# Your WordDictionary object will be instantiated and called as such:
# obj = WordDictionary()
# obj.addWord(word)
# param_2 = obj.search(word)
