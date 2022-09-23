package main

import (
	"fmt"
	"log"
	"sync"
)

type State string

const (
	Follower  State = "follower"
	Candidate State = "candidate"
	Leader    State = "leader"
)

// RequestVote RPC arguments structure. field names must start with capital letters!
type RequestVoteArgs struct {
	// Your data here (2A, 2B).
	Term         int
	CandidateId  int
	LastLogIndex int
	LastLogTerm  int
}

// RequestVote RPC reply structure. field names must start with capital letters!
type RequestVoteReply struct {
	// Your data here (2A).
	Term        int
	VoteGranted bool
}

type Raft struct {
	mu    sync.Mutex
	me    int
	peers []int
	state State

	currentTerm int
	votedFor    int
}

// requestForVotes
func (rf *Raft) AttemptElection() {
	rf.mu.Lock()
	rf.state = Candidate
	rf.currentTerm++
	rf.votedFor = rf.me

	log.Printf("[%d] attempting an election at term %d", rf.me, rf.currentTerm)

	votes := 1
	done := false
	term := rf.currentTerm
	rf.mu.Unlock()

	cond := sync.NewCond(&rf.mu)

	for _, server := range rf.peers {
		if server == rf.me {
			continue
		}
		go func(server int) {
			voteGranted := rf.CallRequestVote(server, term)
			log.Printf("[%d] ?voteGranted=%v)", rf.me, voteGranted)
			if !voteGranted {
				return
			}

			//tally votes
			rf.mu.Lock()
			defer rf.mu.Unlock()

			votes++
			cond.Broadcast()
			log.Printf("[%d] got vote from %d", rf.me, server)
			if done || votes <= len(rf.peers)/2 {
				return
			}
			done = true
			if rf.state != Candidate || rf.currentTerm != term {
				return
			}
			rf.state = Leader
			log.Printf("[%d] We got enough votes, we are now leader (currentTerm=%v state=%v votes=%v)!", rf.me, rf.currentTerm, rf.state, votes)
		}(server)
	}

	rf.mu.Lock()
	for votes <= len(rf.peers)/2 {
		cond.Wait()
	}
	if votes > len(rf.peers)/2 {
		println("won election!", votes, rf.state)
	} else {
		println("lost", rf.state)
	}
	rf.mu.Unlock()
}

// send a RequestVote RPC to a server.
// server is the index of the target server in rf.peers[].
// expects RPC arguments in args.
// fills in *reply with RPC reply, so caller should
// pass &reply.
// the types of the args and reply passed to Call() must be
// the same as the types of the arguments declared in the
// handler function (including whether they are pointers).
//
// The labrpc package simulates a lossy network, in which servers
// may be unreachable, and in which requests and replies may be lost.
// Call() sends a request and waits for a reply. If a reply arrives
// within a timeout interval, Call() returns true; otherwise
// Call() returns false. Thus Call() may not return for a while.
// A false return can be caused by a dead server, a live server that
// can't be reached, a lost request, or a lost reply.
//
// Call() is guaranteed to return (perhaps after a delay) *except* if the
// handler function on the server side does not return.  Thus there
// is no need to implement your own timeouts around Call().
//
// look at the comments in ../labrpc/labrpc.go for more details.
//
// if you're having trouble getting RPC to work, check that you've
// capitalized all field names in structs passed over RPC, and
// that the caller passes the address of the reply struct with &, not
// the struct itself.
func (rf *Raft) CallRequestVote(server int, term int) bool { // sendRequestVote
	//rf.mu.Lock()
	//rf.mu.Unlock()
	log.Printf("[%d] sending request vote to %d", rf.me, server)
	args := RequestVoteArgs{
		//Term:        rf.currentTerm,
		Term:        term,
		CandidateId: rf.me,
	}

	var reply RequestVoteReply
	//ok :=
	rf.sendRequestVote(server, &args, &reply)
	log.Printf("[%d] finish sending request vote to %d", rf.me, server)

	if reply.VoteGranted != true {
		return false
	}
	log.Printf("voteGranted=%v ok=%v", reply.VoteGranted)
	//if !ok {
	//	return false
	//}

	// process the reply

	return true
}

func (rf *Raft) sendRequestVote(server int, args *RequestVoteArgs, reply *RequestVoteReply) {
	log.Printf("[%d] sendRequestVote %d %d %d", server, args.Term, rf.me, args.CandidateId)
	rf.RequestVote(args, reply)
}

//s0.CallRequestVote, acquire the lock
//s0.CallRequestVote, sendRPC to s1
//s1.CallRequestVote, acquire the lock
//s1.CallRequestVote, sendRPC to s0
//s0.Handler, s1.Handler trying to acquire lock
//The solution to the problem is to not hold locks thru RPC call

// RequestVote - Invoked by candidates to gather votes (§5.2).
// RequestVote RPC handler. [HandleRequestVote]
func (rf *Raft) RequestVote(args *RequestVoteArgs, reply *RequestVoteReply) {
	log.Printf("[%d] received request vote from %d", rf.me, args.CandidateId)
	//rf.mu.Lock()
	//defer rf.mu.Lock()

	log.Printf("[%d] handling request vote for %d", rf.me, args.CandidateId)

	if args.Term < rf.currentTerm {
		reply.Term = rf.currentTerm
		reply.VoteGranted = false
		return
	}

	//if args.Term > rf.currentTerm {
	//	rf.stepDownToFollower(args.Term)
	//}

	reply.Term = rf.currentTerm
	reply.VoteGranted = false

	if rf.votedFor < 0 || rf.votedFor == args.CandidateId {
		reply.VoteGranted = true
		rf.votedFor = args.CandidateId
		log.Printf("[%d] VoteGranted votedFor %d", args.CandidateId, rf.votedFor)
		return
	}

	return
}

func main() {
	rf := &Raft{}
	rf.peers = []int{101, 102, 103, 104, 105}
	fmt.Println(rf.peers)
	rf.AttemptElection()
}
