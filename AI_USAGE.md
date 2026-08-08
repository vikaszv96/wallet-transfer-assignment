# AI Usage Disclosure

## 1. Tool

Claude Code (Anthropic), running as an interactive coding agent directly in
my editor.

## 2. How I generally used it

Pair-programming, not delegation. I set the direction (Go, Gin, layered
architecture, no single giant file, must actually run and be testable) and
reviewed/verified everything as it was built rather than accepting output
blindly:

- The design (`DESIGN.md`) was written and reviewed before any code, per the
  assignment's documentation-first workflow.
- I had it actually run the build, the unit tests, the integration tests
  against a real Postgres container, and a live concurrency demo against the
  running server (12 concurrent HTTP requests racing for the same wallet's
  balance) -- not just generate code and claim it works.
- After the implementation was "done," I asked it to walk me through nearly
  every part of the codebase line by line (see the second half of the prompt
  list below) specifically so I could understand and be able to defend every
  design decision myself, not just submit generated code I couldn't explain.
- That review process caught two real issues, which were then fixed: a
  timestamp bug (`createdAt` coming back as the zero value because the DB's
  `now()` default was never read back into the Go struct), and a missing
  business rule (nothing stopped a transfer between wallets of different
  currencies).

## 3. Prompts used (chronological, this repository's session)

Reproduced close to verbatim, typos included, for an accurate record.

### Getting the project set up
1. "ok now i have to work in one project. SO forst clone this repo: https://github.com/Robustrade/wallet-transfer-assignment and import it here in this workspace for me fast"
2. "are you idiot or what? docvid is different project. Undo it and add it in UNTITLES workspace directly simialr to docvid is imprted as new project"
3. "I am unable to click in menu options? i dnt know why" (unrelated IDE issue, resolved separately)
4. "ok now imported. So now frst give me the broad idea what is this assignment about and short summary of what thungs we have to do? analyze the whole project"
5. "means we have to build like money can be tranferred from one wallet o anither wallet or not right?"

### Building the solution
6. "Ok start building it. We have to use Golang only and use the framwork Gin to solve the complexity. and also make it modualr like dont write all code in one file please! it should be accoridng the assigment.md and modualarzed. Begin now and in the last show me the working how it will work and sumamry of it."

### Submission process
7. "How do i hvae to submit this? Do I need to explain this too to them?"
8. Follow-up answers: GitHub username/fork details, branch name `solution/vikasyadav`, choice to push manually rather than have the agent authenticate.

### Understanding the codebase (asked after the implementation existed, to verify understanding)
9. "ok so like in complex go project, do we define all api endpoints in same router.go file?"
10. "What is go.mod and go.sum used for? Aso what is Makefile?"
11. "Compare these files with nodejs typical project: go.mod and go.sum and Makefile"
12. "ok now i have to undersatnd the project. SO here all API endpoints are defined under router folder? handler folder act as controllersand service folder for logic and repository for databse callings and domains for?"
13. "But why its inside cmd/server but not in root project folder?"
14. "ok how this project gets run now in local?"
15. "ok now explain cleanly and slowly the main.go file code"
16. "Why you haven't added the import for internal/domain? And also why you have used robustrade/ ??"
17. "So if in future, if we have more handlers then we will define more args here? func New(transferHandler *handler.TransferHandler, walletHandler *handler.WalletHandler)"
18. "I want to undersatnd what's the use of r here: r := router.New(...) and r := gin.New() inside router.go"
19. "r.POST(\"/transfers\", transferHandler.Create) ... Here in this api how its checked the payload is sent or not? ... I want to underatnd how payload in POST and PUT apis and params and queryparams with GET or PUT or Delete APIs are sent and received."
20. "But its direclty used here without importing? var req CreateTransferRequest"
21. "so that means even transfer_handler can access all structs or funcs from wallet_handlers too?"
22. "ok now explain the full api r.POST(\"/transfers\", transferHandler.Create) how contorllers service and others are called and what logic is being done here means what we are doing with this API"
23. "Ok now just explain me the whole logic how this wallet tranfer will work and when we have to execute what API?"
24. "How its incrementing and decrementing inside wallet while tranfer?"

### Quality check before submission
25. "ok so whatever you have designed or written the code this is production level only right? not fuzzy code. The interviewer panel should get impressed with this." -- this prompt led to a deliberate honest audit rather than reassurance, which surfaced and fixed the cross-currency gap described above.

## AI's role vs. mine

- AI: wrote the design doc and implementation, ran and verified the build/
  tests/integration tests/live demo, explained the codebase in depth on
  request, and performed a self-audit that found real gaps when asked to
  assess production-readiness honestly.
- Me: set the technical constraints and direction, made the fork/branch/
  submission decisions, chose not to grant the agent GitHub push
  credentials (pushing this myself instead), and worked through the
  implementation line by line until I could explain the idempotency
  strategy, the locking strategy, and the request lifecycle myself.
