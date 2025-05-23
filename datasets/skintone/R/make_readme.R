# ==============================================================================
# Title        : make_readme.R
# Description  : Knit readme.Rmd to README.md
# Author       : SAIC IDSL
# Reference    : https://doi.org/10.1145/3730409
# Repository   : https://github.com/TheMdTF/mdtf-public
# License      : https://github.com/TheMdTF/mdtf-public/blob/master/LICENSE.md
# ==============================================================================
root <- this.path::this.dir()
rmd.file <- 'readme.Rmd'
out.file <- '../README.md'

knitr::knit(input = file.path(root, rmd.file),
            output = file.path(root, out.file))
