#include <stdio.h>
#include <unistd.h>

#include "cs107thread.h"

sem_t sem;

// Function executed by the threads
void *print_message(void *ptr)
{
  sem_wait(&global_lock);
  char *message;
  message = (char *)ptr;
  printf("%s \n", message);
  sem_post(&sem);
  return NULL;
}

int main()
{
  sem_init(&sem, 0, 0);

  const char *message1 = "Thread 1";
  const char *message2 = "Thread 2";

  ThreadNew(message1, print_message, (void *)message1);
  ThreadNew(message2, print_message, (void *)message2);

  // sleep(3);
  StartAllThread();
  // wait for both threads to finish
  sem_wait(&sem);
  sem_wait(&sem);

  printf("Both threads finished\n");
  return 0;
}
