include $(GOROOT)/src/Make.$(GOARCH)

myall: tangle all

tangle:
	notangle -R$(TARG).go $(TARG).nw > $(TARG).go

CLEANFILES+= *.aux *.log *.tex *.pdf $(TARG).go
TARG=rpu
GOFILES=\
	$(TARG).go\

include $(GOROOT)/src/Make.cmd

doc:
	noweave -delay -latex $(TARG).nw > $(TARG).tex
	pdflatex $(TARG).tex
