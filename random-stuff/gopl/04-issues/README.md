# Github issue tracker

Track the guthub issues programmatically from command line

## Usage

Here are some live examples

- Query for open issues of simplecv with page 1
```bash
⇒  go run main.go simplecv is:open page:1
INFO[0000] initiating issue search for https://api.github.com/search/issues?q=simplecv+is%3Aopen+page%3A1  function=SearchIssues
INFO[0000] search query response status: 200 OK          function=SearchIssues status="200 OK"
10 issues:
#    1 yanningzuo About SimpleCV installation troubleshoot
#  491    codetab Documentation page: CSS needs to be updated and other b
#   11    Nivasb1 Not able to access my Camera in Ubuntu
#  668 ctrochalak Release 1.4
#  542 jayrambhia OpenCV has stopped supporting Legacy Code
#  339    cversek Expression ```np.array(self.getMatrix())``` for Image o
#  654 eleventhHo Error thrown when the isDone() method is called in the
#   13       twtw remove linux packages
# 4906    guevara Object Detection for Dummies Part 1: Gradient Vector HO
# 1422    alecive [xelatex][windows] Expected PDF file not found
```

- Query for open issues of plantcv
```bash
⇒  go run main.go plantcv is:open
INFO[0000] initiating issue search for https://api.github.com/search/issues?q=plantcv+is%3Aopen  function=SearchIssues
INFO[0000] search query response status: 200 OK          function=SearchIssues status="200 OK"
63 issues:
#  666 HaleySchuh Remaining analysis fxn label updates
#  663 HaleySchuh Add "label" parameter to morph&hyperspectral analysis f
#  667  JimXu1989 Please support the newest version of opencv
#  587 DannieShen Add a sub-package "time series" into PlantCV
#  655 DannieShen Visualize show instances
#  643  nfahlgren Reorganize tests
#  617 dschneider Executable scripts
#    2 AdamDimech Data science tools plantcv issue 437
#  437 AdamDimech PhenoFront and PlantCV
#    1   Mahi-Mai cannot import color_palette
#   36  nfahlgren Update structure for how SPICE is integrated
#  651 bharathred How do i read multiple images in Plantcv?
#  616 dschneider cross-platform compatible command line scripts
#  228   dlebauer Obsolete PlantCV extractors directories to delete
#  656 dschneider substitute deprecated MAINTAINER in Dockerfile
#  657  nfahlgren Move rotate function to transform submodule
#  624 AdamDimech Image Demosaicing Quality in PlantCV versus LemnaBase
#  653 anshikasha Analysis of thermal image of leaf for diseases. Check i
#  661 vtrifunovi Multi-plant parallelization
#  649 dschneider correct readbayer behavior
#  177 TimeScienc Review PlantCV code and report back
#  569 DannieShen pcv.visualize.pseudocolor and pcv.threshold.mask_bad
#  645 dschneider update dockerfile to conda-forge
#  568 DannieShen Add function to threshold method
#    4 vektorious Write installation instruction
#    7 Blackciphe Change the track bar default to 6
#  642 dschneider readbayer color channels are out of order
#  592 SethPolydo Color_clustering_alhorithm
#  639   bmsowder Data file types for Hyperspectral Workflow.
#  203  nfahlgren Hyperspectral submodule
```

- Query for all open issues form plottly
```bash
⇒  go run main.go is:open javascript plotly
INFO[0000] initiating issue search for https://api.github.com/search/issues?q=is%3Aopen+javascript+plotly  function=SearchIssues
INFO[0001] search query response status: 200 OK          function=SearchIssues status="200 OK"
899 issues:
# 3229    jks-liu Proper Plotly JavaScript setup in IJulia
#  113  elong0527 Provide column width in cellInfo object for JavaScript
#   84     DarWik SEO issue: 7 external JavaScript file found.
#   58 nicolaskru Plotly.js 2.0 compatibility
#   17   kvhooreb Make the schema table filterable and sortable
#    2  kylefritz Plotly instead of chartjs
#  362 nicolaskru Plotly.js 2.0 compatibility
# 1911 nicolaskru Plotly.js 2.0 compatibility
#  125    lexvand Uitbreiden features Plotly javascript library
#   34 florian-hu Create embedded javascript networking plot
# 5271 Alex-Monah Question: Will Plotly.js get Plotly Express?
#  333   pyup-bot Update dash to 1.19.0
#   94       dmt0 Relayouting event of plotly.js is not documented
#   76   pyup-bot Pin dash to latest version 1.19.0
#   12 dependabot Bump lxml from 4.2.4 to 4.6.2 in /JavaScript-Plotly
#   15 stephhazli Explore options to tailor plotly ModeBar
#   19   kjetilkl FASTP does not display graphs properly due to CORS rest
#   52    m-rossi Add support for JavaScript-based visualizations
#  114   pyup-bot Scheduled weekly dependency update for week 03
# 5395 nicolaskru v2 changes
# 4364 AnupamSara Graph animations do not work in jupyter notebook
# 5403 frankroesk [Feature Request] New Sizing Option to Repeat Source Im
#50718 chowington plotly.js: Add separate plot data definition for pie tr
# 4921 juliangomm [feature request] Trendlines in Plotly.js
#   13 dependabot Bump ecstatic from 3.3.1 to 4.1.4 in /fabcar/javascript
# 1775     NHDaly Allow plotly plots to auto-resize to fill the window.
#  248   pyup-bot Update dash to 1.19.0
#   74    epykure Add Plotly candlestick charts
#   12 dependabot Bump ini from 1.3.5 to 1.3.8 in /fabcar/javascript/publ
#   27   sean-mcl Add support for plotly events
```