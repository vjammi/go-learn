## Raft Impl
Reference: https://youtu.be/UzzcUS2OHqo?t=2804

So yeah basically I'm going to show you two bugs that we commonly see in people's raft implementations 
there's a lot of bugs that are pretty common but I'm just going to focus on two of them so in this
first example we sort of have a start of a raft implementation for that's sort of
like what you might see for to a just the beginnings of one so in our raft state we have primarily
the current status of the raft pier either follower candidate or leader and we have these two state variables that
were keeping track of the current term and who we voted for in the current term so I'm I want us to focus though on
these two functions attempt election and call request vote so in a temptin we're
just going to set our state to candidate increment our current term vote for
ourselves and then start sending out request votes to all of our raft peers and so this is similar to some of the
patterns that Anish showed where we're going to loop through our peers and then
for each one in a go routines separately call this call request vote function in
order to actually send an RPC to that peer alright so in call request vote we're
going to acquire the lock prepare arguments for our request Ville RPC call
based on by setting it to the current term and then actually perform the RPC
call over here and finally based on the response we will reply back to this this
attempt election function and the attempt election function eventually should tally up the votes to see if it got a majority of the votes and can
become leader so what happens when we run this code so in theory what we might
expect to happen is four so there's going to be some code that's going to spawn a few graph spears and actually
try to attempt elections on them and what should happen are we just start
collecting votes from other peers and then we're not actually going to tally them up but hopefully nothing weird goes wrong
but actually something is going to go wrong here and we actually activated
goes deadlock detector and somehow we ran into a deadlock so let's see what happened for now let's focus on what's
going on with the server zero so server zero it says it starts attempting an
election at term one that's just starting the attempt election function it will acquire the lock set some of the
set some stuff up for performing the election and then unlock then it's going
to send out a request vote RPC - server - it finishes processing that request
vote RPC over here so we're just printing right before and after we actually send out the RPC and then it
sends out a request vote RPC - server one but after that it never we never actually see it finish sending the
request vote RPC so it's actually stuck in this function call waiting for the
RPC response from server 1 all right now let's look at what's everyone's doing so it's it's pretty much the same thing it
sends a request vote I received a server to that that succeeds it finishes processing that request vote the
response from server 2 then it sends this RPC to zero and now what's actually
happening is 0 & 1 are sort of waiting for the RPC responses from each other they both sent out an RPC call but not
yet got the response yet and that's actually sort of the cause of our deadlock so really what's the reason
that we're dead locking is because we're holding this lock through our RPC calls over here in the core requests vote
function we acquire our mutex associated with our raft peer and we only unlock at
the end of this function so throughout this entire function we're holding the lock including when we try to contact
our peer to get the vote and later when
we handle this request vote RPC we
actually only see it at the beginning of this function in the handler we're also trying to acquire the lock but we never
actually succeed in acquiring the lock so just to make this a little bit more clear the the sort of order of
operations is happening is in call request vote
server zero is first going to acquire the lock and send an RPC call to server
one and then simultaneously and separately server one is going to do the
same thing it's going to enter its call request vote function acquire the lock and send this RPC call to server zero
now in server zeros handler and server ones handler they're trying to acquire
the lock but they can't because they already are acquiring the lock and trying to send the RPC call to each
other and that that's actually what's leading to the deadlock situation so to
solve this basically we want you to not hold locks through RPC calls and that's the solution to this problem in fact we
don't need the lock here at all instead of trying to read the current term when we enter this call request vote function
we can pass this as an argument here save the term when we had acquired the
lock earlier in this attempt election and just passed this as a as a variable to call request vote so that actually
removes the need to acquire the lock at all in call request vote alternatively
we could lock while we're preparing the arguments and then unlock before actually performing the call and then if
we need to to process the reply we could lock again afterwards so it's just make sure to unlock before making it
obviously call and then if you need to you can acquire the lock again so now if
I save this then so it's still
activating the deadlock detector but that's actually just because we're not doing anything at the end but now it's
actually working we finished sending the request votes on both sides and all the operations that we wanted to complete are complete all
right any questions about this example
yeah so not it's sort of so you might need to use locks when you are preparing
the arguments or processing the response but yeah you shouldn't hold a lock through the RPC call while you're
waiting for the other peer to respond and there's actually another reason to that in addition to deadlock the other
problem is that in some tests we're going to sort of have this unreliable
network that could delay some of your RPC messages potentially by like 50 milliseconds and in that case if you
hold the lock through an RPC call then any other operation that you try to do during that 50 milliseconds won't be
able to complete until that RPC response is received so that that's another issue that you might run into if you hold the
long so it's both to make things more efficient and to avoid these potential deadlock situations
all right so just one more example this is again using a similar draft
implementation so again in our raft state we're going to be keeping track of whether a fuller candidate leader and then also these two state variables in
this example I want you to focus on this attempt election function so now we've first implemented the change that I just
showed you to store the term here and pass it as a variable to our function
that collects the request votes but additionally we've implemented some functionality to add up the votes so
what we'll do is we'll create a local variable to count the votes and whenever
we get a vote if the vote was not granted we'll return immediately from this go routine where we're processing the boat
otherwise we'll acquire the lock before editing this shared local variable to
count up the votes and then if we did not get a majority of the votes will return immediately otherwise we'll make
ourselves the leader so as with the other example I mean initially if you
look at this if I look at this like it seems reasonable but let's see if anything can go wrong all right so this
is the log output from one run and one thing you might notice is that we've actually elected two leaders on the same
term so server zero it was elected made itself a leader on
term two and server one did as well it's okay to have a leader elected on different terms but here where we have
one on the same term that that should never happen alright so how did this actually come up so let's start from the
top so at the beginning server zero actually attempted an election at term one not turn two and it got its votes
from both of the other peers but for whatever reason perhaps because those reply messages from those peers were
delayed it didn't actually process its process those votes until later and in
between receiving it like in between attempting the election and finishing the election server one also decided to
attempt an election perhaps because because of server zero was delayed so much server one might
actually ran into the election timeout and then started its own election and it started it on term 2 because it couldn't
have been termed 1 because it already voted for server 0 on on term 1 over here
okay so then server 1 sends out its own request votes 2 servers 2 and 0 at term
2 and now we see that server two votes for server 1 that's fine but server 0 also votes for server 1 this is actually
also fine because server one is asking server 0 for a vote on a higher term and
so what server 0 should do is if you remember from the spec it should set its
current term to that term in the request for RPC message to term 2 and also revert itself to a follower instead of a
candidate alright finally so the real problem is that on this line where
server 0 although it really got enough votes on term 1 it made itself a leader on term - so the reason so one
explanation for why this is happening is because in between where we set up the election our attempt for the election
and where we actually process the votes some other things are happening input in
this case we're actually voting for someone else in between and so we're no
longer on term 1 where we thought we started the election we're now on term 2 and so we just need a double check that
because we don't have the lock while we're performing the RPC calls which is important for its own reasons now some
things might have changed and we need to double check that what we assume is true when we're setting ourselves to the
leader is still true so one way to solve this that there's a few different ways
like to solve this like you could imagine not voting for others while we're in the middle of attempting an election but in this case the simplest
way to solve this at least in this implementation is to just double check
that we're still on the same term and we're still a candidate we haven't reverted to a follower so actually one
thing I want to show you is if we do print out our state over here then we do
see that server 0 became a follower but it's still setting itself to a leader on this line