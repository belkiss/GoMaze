include $(GOROOT)/src/Make.$(GOARCH)

TARG=gomaze
GOFILES=\
	main.go

include $(GOROOT)/src/Make.cmd
