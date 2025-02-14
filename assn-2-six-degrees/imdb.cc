#include <sys/stat.h>
#include <sys/types.h>
// #include <sys/mman.h>
#include <fcntl.h>
#include <unistd.h>

#include "imdb.h"
using namespace std;

const char *const imdb::kActorFileName = "actordata";
const char *const imdb::kMovieFileName = "moviedata";

imdb::imdb(const string& directory)
{
  const string actorFileName = directory + "/" + kActorFileName;
  const string movieFileName = directory + "/" + kMovieFileName;

  actorFile = acquireFileMap(actorFileName, actorInfo);
  movieFile = acquireFileMap(movieFileName, movieInfo);
}

bool imdb::good() const
{
  return !((actorInfo.fd == NULL) || (movieInfo.fd == NULL));
}

// you should be implementing these two methods right here...
bool imdb::getCredits(const string& player, vector<film>& films) const { return false; }
bool imdb::getCast(const film& movie, vector<string>& players) const { return false; }

imdb::~imdb()
{
  releaseFileMap(actorInfo);
  releaseFileMap(movieInfo);
}

// ignore everything below... it's all UNIXy stuff in place to make a file look like
// an array of bytes in RAM..
const void *imdb::acquireFileMap(const string& fileName, struct fileInfo& info)
{
  struct stat stats;
  stat(fileName.c_str(), &stats);
  info.fileSize = stats.st_size;
  info.fd = fopen(fileName.c_str(), O_RDONLY);
  char *ptr = (char *)calloc(info.fileSize, sizeof(char));
  fread(ptr, sizeof(char *), info.fileSize, info.fd);
  return (void *)ptr;
}

void imdb::releaseFileMap(struct fileInfo& info)
{
  if (info.fileMap != NULL) {
    free((char *)info.fileMap);
    info.fileMap = NULL;
  }
  if (info.fd != NULL) fclose(info.fd);
}
