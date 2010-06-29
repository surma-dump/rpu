include $(GOROOT)/src/Make.$(GOARCH)

TARG=rpuemu
GOFILES=\
	$(TARG).go\

include $(GOROOT)/src/Make.cmd
