#!/usr/bin/env bash
#transform by deleting the space, grep and then put the space back
tr -d ' ' | grep -E '00|11|22|33|44|55|66|77|88|99' | sed -r 's/..../& /g'