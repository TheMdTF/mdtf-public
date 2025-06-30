# ==============================================================================
# Title        : colorimetric_skint_tone_scale.R
# Description  : Functions for creating and rendering the CST
# Author       : SAIC IDSL
# Reference    : https://doi.org/10.1145/3730409
# Repository   : https://github.com/TheMdTF/mdtf-public
# License      : https://github.com/TheMdTF/mdtf-public/blob/master/LICENSE.md
# ==============================================================================

# ---- Load Libraries ----
library(ggplot2)
library(dplyr)

# ---- Define Functions ----
make_scale <- function(skin.tone, minL = 20, maxL = 70, nbins = 9) {
  # model a as a function of L
  # model b as a function of L
  # split L into even bins
  # predict a and b from L
  model.C <- lm(data = skin.tone %>% 
                  mutate(l2 = l^2, 
                         C = sqrt(a^2+b^2),
                         h = atan(b/a)), formula = C ~ l + l2)
  model.h <- lm(data = skin.tone %>% 
                  mutate(l2 = l^2, 
                         C = sqrt(a^2+b^2),
                         h = atan(b/a)), formula = h ~ l + l2)
  
  # create the scale
  dL <- (maxL-minL)/nbins
  Lctr <- rev(seq(minL,maxL,by=dL))
  
  data.frame(l = Lctr,
             C = predict(model.C, newdata = data.frame(l = Lctr, l2 = Lctr^2)),
             h = predict(model.h, newdata = data.frame(l = Lctr, l2 = Lctr^2))) %>%
    mutate(a = C * cos(h),
           b = C * sin(h)) %>%
    rowwise() %>%
    mutate(Hex = schemr::lab_to_hex(data.frame(l = l, a = a, b = b))) %>%
    ungroup() %>%
    mutate(
      id = row_number(),
      l = round(l, 10),
      a = round(a, 10),
      b = round(b, 10)) %>%
    select(id, l, a, b, Hex)
}

plot_scale <- function(scale, background = 'white', hex = F) {
  
  scale <- scale %>%
    mutate(hex_label = paste(Hex, 
                             paste0('(', paste(round(l),round(a),round(b), 
                                               sep = ','), ')'), 
                             sep = '\n'))
  
  plt <- ggplot(scale %>%
           mutate(id = row_number()) %>%
           select(id, Hex, hex_label)) +
    geom_rect(aes(xmin = id-.5, xmax=id+0.5, fill = Hex), ymin = -0.5, ymax = 1) + 
    geom_text(aes(x = id, label = id, color = '#000000'), y = -0.75, hjust = 0.5, vjust = 0.5) + 
    scale_fill_identity() +
    scale_color_identity() +
    scale_y_continuous(limits = c(-1,1), expand = c(0,0)) +
    scale_x_continuous(limits = c(0.5,10.5), expand = c(0,0)) +
    annotate(geom = 'text', x = 6.5, y = -.5, vjust = 0, hjust = 0, 
             label = 'https://doi.org/10.1145/3730409') +
    theme_void() +
    theme(plot.margin = margin(t = 0,r = 0,b = 0,l = 0),
          plot.background = element_rect(fill = background))
  
  if (hex) {
    plt <- plt + geom_text(aes(x = id, label = hex_label, color = 'white', 
                               y = 0.25, vjust = 0.5), size = 8/.pt)
  }
  
  plt
  
}

write_commented_csv <- function(data, file, comments = NULL, ...) {
  con <- file(file, open = "wt")
  on.exit(close(con))
  
  if (!is.null(comments)) {
    comment_lines <- paste0("# ", comments)
    writeLines(comment_lines, con)
  }
  
  write.csv(data, con, ...)
}


# ---- Load Data and Create Scale ----

root <- this.path::this.dir()
skin.tone.data <- '../data/idsl_skin_tone_sample.csv'
scale.out.csv <- '../data/colorimetric_scale.csv'
scale.png <- '../images/scale.png'
gamut.png <- '../images/gamut.png'

skin.tone <- read.csv(file.path(root, skin.tone.data), 
                      stringsAsFactors = F, comment.char = '#')

scale <- make_scale(skin.tone)

header <- c("==============================================================================",
            "Title        : colorimetric_scale.csv",
            "Description  : Colorimetric Skin Tone Scale",
            "Author       : SAIC IDSL",
            "Reference    : https://doi.org/10.1145/3730409",
            "Repository   : https://github.com/TheMdTF/mdtf-public",
            "License      : https://github.com/TheMdTF/mdtf-public/blob/master/LICENSE.md",
            "==============================================================================")

write_commented_csv(file = file.path(root, scale.out.csv),
                    comments = header,
                    data = scale, row.names = F)

# ---- Plot Skin Tone Scale ----

plt.scale <- plot_scale(scale, hex = T)
ggsave(plt.scale, filename = file.path(root, scale.png), 
       width = 6, height = 2, units = 'in')

# ---- Plot Skin Tone Gamut ----

gamut <- skin.tone %>%
  rowwise() %>%
  mutate(Hex = schemr::lab_to_hex(data.frame(l=l,a=a,b=b))) %>%
  ungroup() %>%
  mutate(h = (atan2(b,a)/pi*180 + 360) %% 360, 
         C = sqrt(a^2+b^2)) %>%
  filter(h < 80) # remove one point that is out of the visualization scale

scale.plt <- scale %>%
  mutate(h = (atan2(b,a)/pi*180 + 360) %% 360, 
         C = sqrt(a^2+b^2))

box <- gamut %>%
  summarise(max.h = round(quantile(h,.999)),
            min.h = round(quantile(h,.001)),
            max.L = round(quantile(l,.999)),
            min.L = round(quantile(l,.001))) 

y.off = -5
x.off = 15

plt <- ggplot(gamut) + 
  geom_point(aes(y=h,x=l,color=Hex), size = 12, show.legend = F) + 
  scale_color_identity() +
  annotate(geom='segment',x=box$min.L, xend=box$max.L, y=y.off, yend=y.off, color = 'white') + 
  annotate(geom='segment',x=box$min.L, xend=box$min.L, y=y.off, yend=y.off-1, color = 'white') + 
  annotate(geom='segment',x=box$max.L, xend=box$max.L, y=y.off, yend=y.off-1, color = 'white') + 
  annotate(geom='text', x = box$min.L, y = y.off-4, label = box$min.L, color = 'white') +
  annotate(geom='text', x = box$max.L, y = y.off-4, label = box$max.L, color = 'white') +
  annotate(geom='text', x = box$min.L + 0.5*(box$max.L-box$min.L), y = y.off-4, label = 'Lightness', color = 'white') +
  annotate(geom='segment',x=x.off, xend=x.off, y=box$min.h, yend=box$max.h, color = 'white') + 
  annotate(geom='segment',x=x.off-0.3, xend=x.off, y=box$min.h, yend=box$min.h, color = 'white') + 
  annotate(geom='segment',x=x.off-0.3, xend=x.off, y=box$max.h, yend=box$max.h, color = 'white') + 
  annotate(geom='text', x = x.off-1, y = box$min.h, label = box$min.h, color = 'white') +
  annotate(geom='text', x = x.off-1, y = box$max.h, label = box$max.h, color = 'white') +
  annotate(geom='text', x = x.off-1, y = box$min.h + 0.5*(box$max.h-box$min.h), label = 'Hue (deg)', color = 'white', angle = 90) +
  geom_text(aes(y=h,x=l,color='white', label=id), data = scale.plt) +
  annotate(geom = 'text', x = box$min.L, y = box$min.h, 
           vjust = 0, hjust = 0, label = 'https://doi.org/10.1145/3730409') +
  scale_y_continuous(limits = c(-20,80), breaks = c(seq(0,300,60))) +
  xlim(10,70) +
  theme_void() +
  theme(plot.background = element_rect(fill='grey40'),
        panel.background = element_rect(fill='grey40'))

ggsave(plt, filename = file.path(root, gamut.png),
       width = 6, height = 4, units = 'in')
