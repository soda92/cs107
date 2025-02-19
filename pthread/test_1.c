#include <stdio.h>

#include "pthread.h"
#include "semaphore.h"

sem_t sem;

// Function executed by the threads
void *print_message(void *ptr)
{
  char *message;
  message = (char *)ptr;
  printf("%s \n", message);
  sem_post(&sem);
  pthread_exit(NULL);
}

int main()
{

  sem_init(&sem, 0, 0);
  pthread_t thread1, thread2;
  char *message1 = "Thread 1";
  char *message2 = "Thread 2";
  int ret1, ret2;

  // Create thread 1
  ret1 = pthread_create(&thread1, NULL, print_message, (void *)message1);
  if (ret1) {
    fprintf(stderr, "Error creating thread 1\n");
    return 1;
  }

  // Create thread 2
  ret2 = pthread_create(&thread2, NULL, print_message, (void *)message2);
  if (ret2) {
    fprintf(stderr, "Error creating thread 2\n");
    return 1;
  }

  // Wait for threads to finish
  pthread_join(thread1, NULL);
  pthread_join(thread2, NULL);
  sem_wait(&sem);
  sem_wait(&sem);

  printf("Both threads finished\n");
  return 0;
}