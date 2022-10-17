package main

/**
At a high level you grab the lock, you mutate the shared data and you unlock.
Turns out that that is a useful starting point but not the complete story

Below is a back data where you might see violations


what we intended here was for this decrement and increment to happen atomically
but instead of what we ended up writing was code that decrement atomically and then increments atomically
and so in this particular code actually like we won't lose money in the long term like if we let these
threads run and then wait till they finish and then check the total it will indeed be what it started out as but
while these are running since this entire block of code is not atomic we can temporarily observe these violations

and so at a higher level the way should think about locking is not just like locks are to protect access to shared data
but locks are meant to protect invariants
you have some shared data that multiple people might access and there's some properties that hold on that shared data
like for example here I is the programmer decided that I want this property that alice + Bob should equal some constant
and that should always be that way I want that property to hold

but then it may be the case that different threads running concurrently are making changes to this data and
might temporarily break this invariant here, like here when I decrement from Alice, temporarily the sum Alice Plus Bob has changed
but then this thread eventually ends up restoring this invariant here and so locks are meant to protect invariants

At a high level you grab a lock then you do some work that might temporarily break the invariant but then you restore
the invariant before you release the lock so nobody can observe these in progress updates
and so the correct way to write this code is to actually have less use of lock and unlock
we have lock then we do a bunch of work and then we unlock

and now when we run this code we see no more printouts like this that we never have this audit thread observe that the total is not what it should be all right so that's the right
way to think about locking at kind of a high level you can think about it as
makenew sure you grab locks when every access shared data like that is a rule but another important rule is locks
protect invariants - so grab a lock manipulate things in a way that might break the invariants but restore them
afterwards and then release the lock

Another way you can think about it is locks can makenew regions of code atomic not just like single statements or single updates to shared variables
*/
