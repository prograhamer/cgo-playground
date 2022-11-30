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
   return tree;
}

int HETreeAdd(HETree *tree, int value) {
   HENode *node = NULL;

   if (tree == NULL) {
      errno = EINVAL;
      return -1;
   }

   if (tree->root == NULL) {
      node = malloc(sizeof(HENode));
      if (node == NULL) {
         return -1;
      }
      tree->root = node;
   } else {
      HENode *this = tree->root;
      HENode **next = NULL;
      while (node == NULL) {
         if (value >= this->value) {
            next = &(this->right);
         } else {
            next = &(this->left);
         }

         if (*next == NULL) {
            node = malloc(sizeof(HENode));
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
