package main

import "fmt"

type connection struct {
	movie  film
	player string
}

/**
 * Convenience Class: path
 * -----------------------
 * The path class is a container class designed to
 * to store a sequence of actors/actresses and movies.
 * It is by no means intelligent, as it'll allow any
 * sequence of movies and players and provide none
 * of the consistency checks one might want.  You're
 * free to change this code to include those consistency
 * checks, or you may leave it alone and use it as is.
 */
type path struct {
	startPlayer string
	links       []connection
}

func (p *path) print() {
	if len(p.links) == 0 {
		fmt.Println("[Empty path]")
		return
	}

	fmt.Print("\t%s was in", p.startPlayer)

	for i := 0; i < len(p.links); i++ {
		fmt.Println(
			"\"%s\" (%s) with %s.",
			p.links[i].movie.title, p.links[i].movie.year, p.links[i].player)
		if i+1 == len(p.links) {
			break
		}
		fmt.Print("\t%s was in ", p.links[i].player)
	}
}

/**
 * Constructor: path
 * -----------------
 * Initializes a path to contain the specified player
 * and no one else.  The path grows because the client
 * appends movie/actor pairs, and the path shrinks when
 * the client calls undoConnection.
 */
func NewPath(startPlayer string) path {
	/**
	 * All paths are designed to store the path from a specified
	 * actor or actress to another actor or actress through a series
	 * of movie-player connections.  A new path is always a partial
	 * path because it only knows of the first player in the chain.
	 * As a result, the embedded vector should be set to be empty, because
	 * each entry in the vector is one leg in the path from an actor to
	 * another.
	 */
	var p path
	p.startPlayer = startPlayer
	return p
	// ommission of links from init list calls the default constructor
}

/**
 * Method: getLength
 * -----------------
 * Replies with the number of movies in place to
 * form the path from the first player to the last.
 *
 * @return the number of movies currently making up the
 *         path.
 */
func (p *path) getLength() int {
	return len(p.links)
}

/**
 * Method: addConnection
 * ---------------------
 * Blindly adds the specified movie and actor to the
 * path.  No integrity checking is done, so it's the
 * responsibility of the client to make sure that the
 * specified player really did appear in the movie, and it's
 * also the client's responsibility to ensure that the last
 * player prior to the addConnection message was also in the
 * movie.  In theory, there is no limit to the number of
 * of movie-player connections that can be added.
 *
 * The implementation makes a deep copy of the specified movie and
 * actor, so you needn't worry about the memory management issues that
 * come up here.
 *
 * @param movie a reference to the film record starring both the specified
 *              player and the last player in the path.
 * @param player a reference to the actor/actress appearing in the specified film.
 */
func (p *path) addConnection(movie film, player string) {
	/**
	 * Simply tack on a new connection pair to the end of the links vector.
	 * It ain't our business to be checking for consistency of connection, as
	 * that's the resposibility of the surrounding class to decide (or at
	 * least we're making it their business.
	 */
	var c connection
	c.movie = movie
	c.player = player
	p.links = append(p.links, c)
}

/**
 * Method: undoConnection
 * ----------------------
 * Pulls the last movie-player connection off the path, unless
 * the length of the path (as defined by our getLength() method) is 0,
 * in which case nothing happens.
 */
func (p *path) undoConnection() {
	if len(p.links) == 0 {
		return
	}
	p.links = p.links[:len(p.links)-1]
}

/**
 * Method: getLastPlayer
 * ---------------------
 * Returns the address of the last player in the
 * path.  Self-explanatory.
 *
 * @return the address of a string storing the name of the
 *         last player.
 */
func (p *path) getLastPlayer() string {
	if len(p.links) == 0 {
		return p.startPlayer
	}
	return p.links[len(p.links)].player
}

/**
 * Method: reverse
 * ---------------
 * Reverses the receiving path.
 */
func (p *path) reverse() {
	reverseOfPath := NewPath(p.getLastPlayer())

	for i := len(p.links) - 1; i > 0; i-- {
		reverseOfPath.addConnection(
			p.links[i].movie, p.links[i-1].player)
	}
	if len(p.links) > 0 {
		reverseOfPath.addConnection(
			p.links[0].movie, p.startPlayer)
	}
	p = &reverseOfPath
}
