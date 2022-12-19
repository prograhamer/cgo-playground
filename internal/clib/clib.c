#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <errno.h>

#include "clib.h"

void _he_tree_walk_node(HENode *node, void (*cb)(int));

char *Reverse(const char *data, size_t len) {
   char *target;
   target = (char *) malloc(len);

   for (int i = 0; i < len; i++) {
      target[i] = data[len - i - 1];
   }

   return target;
}

void ReverseInPlace(char *data, size_t len) {
   char t;
   size_t mid = len / 2;

   for (int i = 0; i < mid; i++) {
      t = data[i];
      data[i] = data[len-i-1];
      data[len-i-1] = t;
   }
}

HETree *HETreeInit() {
   HETree *tree = (HETree*)malloc(sizeof(HETree));
   memset(tree, 0, sizeof(HETree));
   return tree;
}

int _he_tree_free_node(HENode *node) {
   if (node == NULL) {
      return 0;
   }

   int freed = 0;

   if (node->left != NULL) {
      freed += _he_tree_free_node(node->left);
   }
   if (node->right != NULL) {
      freed += _he_tree_free_node(node->right);
   }

   free(node);
   freed += 1;

   return freed;
}

int HETreeFree(HETree *tree) {
   int freed = _he_tree_free_node(tree->root);
   if (freed != tree->size) {
      printf("assertion failed: freed nodes (%d) != tree size (%d)\n", freed, tree->size);
   }
   free(tree);
   return freed;
}

HENode *HENodeInit() {
   HENode *node = (HENode*)malloc(sizeof(HENode));
   memset(node, 0, sizeof(HENode));
   return node;
}

int HETreeAdd(HETree *tree, int value) {
   HENode *node = NULL;

   if (tree == NULL) {
      errno = EINVAL;
      return -1;
   }

   if (tree->root == NULL) {
      node = HENodeInit();
      if (node == NULL) {
         return -1;
      }
      tree->root = node;
   } else {
      HENode *this = tree->root;
      HENode **next = NULL;
      while (node == NULL) {
         if (value > this->value) {
            next = &(this->right);
         } else {
            next = &(this->left);
         }

         if (*next == NULL) {
            node = HENodeInit();
            if (node == NULL) {
               return -1;
            }
            *next = node;
         } else {
            this = *next;
         }
      }
   }

   node->value = value;
   tree->size++;

   return 0;
}

void _he_tree_print_cb(int value) {
   printf("value: %d\n", value);
}

void HETreePrint(HETree *tree) {
   HETreeWalk(tree, _he_tree_print_cb);
}

void HETreeWalk(HETree *tree, void (*cb)(int)) {
   if (tree == NULL || cb == NULL) {
      return;
   }

   _he_tree_walk_node(tree->root, cb);
}

void _he_tree_walk_node(HENode *node, void (*cb)(int)) {
   if (node == NULL) {
      return;
   }

   if (node->left != NULL) {
      _he_tree_walk_node(node->left, cb);
   }

   cb(node->value);

   if (node->right != NULL) {
      _he_tree_walk_node(node->right, cb);
   }
}

void _he_tree_sort_walk_node(HENode *node, int *dest, int *index, int limit) {
   if (node == NULL || index == NULL || *index >= limit) {
      return;
   }

   if (node->left != NULL) {
      _he_tree_sort_walk_node(node->left, dest, index, limit);
   }

   if (*index >= limit) {
      return;
   }

   dest[*index] = node->value;
   (*index)++;

   if (node->right != NULL && *index < limit) {
      _he_tree_sort_walk_node(node->right, dest, index, limit);
   }
}

int *HETreeSort(HETree *tree) {
   if (tree == NULL || tree->size == 0) {
      return NULL;
   }

   int *result = (int *)malloc(tree->size * sizeof(int));
   int index = 0;

   _he_tree_sort_walk_node(tree->root, result, &index, tree->size);

   return result;
}

int HETreeSortNoMalloc(HETree *tree, int *result, int capacity) {
   if (tree == NULL || tree->size == 0 || result == NULL || capacity == 0) {
      return 0;
   }

   int index = 0;

   _he_tree_sort_walk_node(tree->root, result, &index, capacity);

   return index;
}
