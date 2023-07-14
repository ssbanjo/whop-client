release:
	GOPROXY=proxy.golang.org go list -m github.com/ssbanjo/whop-client@v$(version)
