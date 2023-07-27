# Table: hrobot_order_servers

This table shows data for Hrobot Order Servers.

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|description|StringArray|
|traffic|String|
|dist|StringArray|
|arch|IntArray|
|lang|StringArray|
|location|StringArray|
|prices|JSON|
|orderable_addons|JSON|