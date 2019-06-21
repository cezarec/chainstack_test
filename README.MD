# Exporter from Geth node to Prometheus for Grafana dashboard + dashboard itself


Prometheus exporter of syncing metrics from ***Geth*** etherum node.  
**main.go** is a source for metric exporter from the node to Prometheus server.  
Config paramters  are in `const` section of the source itself
The working scheme is next. 
Exporter connets to the Geth node via Go RPC connection and requests metrics every polling interval. After that Prometheus server config includes point to read and store metrics (`scrapping target`) in Prometheus itself. Grafana uses Prometheus as source of data for Dashboards.

As far as all analaysed metrics are measurable describing sync proccess, we use time-interval graphs for all metrics.
Three panels are using for convience.



##  Metrics to analyse:

 - Number if connected peers
 - Number of current synced Block
 - Number of Highest Block in Ethereium Network
 - Quantity of pulled states
 - Quantity of known states




Grafana dashboard link https://grafana.com/dashboards/10240/

***

Tested at:
* Debian 4.9.0.9
* Geth amd64-1.8.27 
* Golang 11.9
* Prometheus 2.9.2
* Grafana 6.20.0
