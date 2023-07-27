# Table: hrobot_order_market_servers

This table shows data for Hrobot Order Market Servers.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|name|String|
|description|StringArray|
|traffic|String|
|dist|StringArray|
|arch|IntArray|
|lang|StringArray|
|cpu|String|
|cpu_benchmark|Int|
|memory_size|Int|
|hdd_size|Int|
|hdd_text|String|
|hdd_count|Int|
|datacenter|String|
|network_speed|String|
|price|String|
|price_setup|String|
|price_vat|String|
|price_setup_vat|String|
|fixed_price|Bool|
|next_reduce|Int|
|next_reduce_date|String|