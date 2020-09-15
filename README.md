# goff
A utility to read space delimited flat file "database" records.

## Overview

I worked somewhere once, that used a (fairly modern) WMS that interfaced with a legacy host OES. The
interface was implemented using space delimited files. Each line of the file was essentially a
record to be inserted into either database. Each field of the row was delimited by the size of the
database column.

### Motivation
There came a time when the host system had to be retired. Mainly because the guy who wrote the thing
had long retired, and developers were really hard to come by.

## goff -- A tool for verifying the interface
