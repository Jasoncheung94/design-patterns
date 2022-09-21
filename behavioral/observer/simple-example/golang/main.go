package main

import "fmt"

type Subject interface {
	subscribe(observer Observer)
	unsubscribe(observer Observer)
	notify()
}

type Observer interface {
	update(string)
	getID() string
}

type NewsFeed struct {
	observers []Observer
	name      string
}

func NewNewsFeed(name string) *NewsFeed {
	return &NewsFeed{
		name: name,
	}
}

func (n *NewsFeed) subscribe(observer Observer) {
	fmt.Println("Subscribing user", observer.getID(), "to feed", n.name)
	n.observers = append(n.observers, observer)
}

func (n *NewsFeed) unsubscribe(observer Observer) {
	for i, o := range n.observers {
		if o.getID() == observer.getID() {
			fmt.Println("Unsubscribing user", observer.getID(), "from feed", n.name)
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
		}
	}
}

func (n *NewsFeed) notify() {
	for _, o := range n.observers {
		o.update(n.name)
	}
}

type User struct {
	id string
}

func NewUser(id string) *User {
	return &User{
		id: id,
	}
}

func (u *User) update(name string) {
	fmt.Printf("User %s received update from %s\n", u.id, name)
}

func (u *User) getID() string {
	return u.id
}

func main() {
	csNewsFeed := NewNewsFeed("Computer Science") // Subject
	engNewsFeed := NewNewsFeed("Engineering")     // Subject
	user1 := NewUser("1")                         // Observer
	user2 := NewUser("2")                         // Observer

	csNewsFeed.subscribe(user1)
	csNewsFeed.subscribe(user2)
	csNewsFeed.unsubscribe(user1)
	engNewsFeed.subscribe(user1)

	csNewsFeed.notify()
	engNewsFeed.notify()
}
