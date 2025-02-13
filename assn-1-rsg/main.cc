/**
 * File: rsg.cc
 * ------------
 * Provides the implementation of the full RSG application, which
 * relies on the services of the built-in string, ifstream, vector,
 * and map classes as well as the custom Production and Definition
 * classes provided with the assignment.
 */

extern "C" {
#include "main.h"
}

#include <fstream>
#include <map>

#include "definition.h"
#include "production.h"
using namespace std;

/**
 * Takes a reference to a legitimate infile (one that's been set up
 * to layer over a file) and populates the grammar map with the
 * collection of definitions that are spelled out in the referenced
 * file.  The function is written under the assumption that the
 * referenced data file is really a grammar file that's properly
 * formatted.  You may assume that all grammars are in fact properly
 * formatted.
 *
 * @param infile a valid reference to a flat text file storing the grammar.
 * @param grammar a reference to the STL map, which maps nonterminal strings
 *                to their definitions.
 */

static void readGrammar(ifstream& infile, map<string, Definition>& grammar)
{
  while (true) {
    string uselessText;
    getline(infile, uselessText, '{');
    if (infile.eof()) return; // true? we encountered EOF before we saw a '{': no more productions!
    infile.putback('{');
    Definition def(infile);
    grammar[def.getNonterminal()] = def;
  }
}

/**
 * Performs the rudimentary error checking needed to confirm that
 * the client provided a grammar file.  It then continues to
 * open the file, read the grammar into a map<string, Definition>,
 * and then print out the total number of Definitions that were read
 * in.  You're to update and decompose the main function to print
 * three randomly generated sentences, as illustrated by the sample
 * application.
 *
 * @param file The grammar file
 */

char *rsg_main(char *file)
{
  ifstream grammarFile(file);
  char *buffer = (char *)calloc(256, sizeof(char));
  if (grammarFile.fail()) {
    cerr << "Failed to open the file named \"" << file << "\".  Check to ensure the file exists. " << endl;
    // return 2; // each bad thing has its own bad return value
    return buffer;
  }

  // things are looking good...
  map<string, Definition> grammar;
  readGrammar(grammarFile, grammar);
  sprintf(buffer, "The grammar file called %s contains %d definitions.", file, grammar.size());
  return buffer;
}
