#pragma once
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

#include "pthread.h"
#include "semaphore.h"

static sem_t global_lock;
static char **thread_names;
static pthread_t tids[100];
static int name_pos = 0;
static bool inited = false;

inline static void
init()
{
  if (!inited) {
    sem_init(&global_lock, 0, 0);
    thread_names = (char **)malloc(100 * sizeof(char *));
    inited = true;
  }
}

// typedef struct func_args
// {
//   void *realfunc;
//   int num_args;
// } FUNC_ARG;

// void myFunc(FUNC_ARG arg)
// {

// }

inline static void
ThreadNew(const char *name, void *(PTW32_CDECL *function)(void *), void *arg)
{
  init();
  char *s = strdup(name);
  thread_names[name_pos++] = s;
  pthread_create(&tids[name_pos - 1], NULL, function, arg);
}

inline static void
StartAllThread()
{
  for (int i = 0; i < name_pos; i++) {
    sem_post(&global_lock);
  }
}
