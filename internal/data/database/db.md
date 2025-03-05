数据库结构:
shortcode表
"expired_time" 时间，和gorm.model.create_time的类型一样
"is_custom"布尔
"short_code" 字符串
"original_url" 字符串


ip访问表：

ip_address ip
user_agent 标识客户端 最近的一次访问的
url 请求地址
access_number 访问次数
访问时间，这个用gorm的model里的创建时间就行了