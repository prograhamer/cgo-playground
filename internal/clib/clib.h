#ifndef __HE_CLIB_H_
#define __HE_CLIB_H_

extern char *Reverse(const char *, size_t);
extern void ReverseInPlace(char *, size_t);

typedef struct _he_tree_node {
   struct _he_tree_node *left;
   struct _he_tree_node *right;
   int value;
} HENode;

typedef struct _he_tree {
   HENode *root;
   int size;
   int depth;
} HETree;

extern HETree *HETreeInit();
extern int HETreeAdd(HETree *, int);
extern void HETreeWalk(HETree *, void (*f)(int));
extern void HETreePrint(HETree *);

#endif
