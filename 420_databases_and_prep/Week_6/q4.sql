select 
	c.cus_code as custcode,
	t.tm_month as month,
	p.p_code as productcode,
	sum(s.sale_units*s.sale_price) as sales
from dwcustomer as c
	inner join dwdaysalesfact as s on c.cus_code=s.cus_code
	inner join dwtime as t on s.tm_id=t.tm_id
	inner join dwproduct as p on s.p_code=p.p_code
group by c.cus_code, t.tm_month, p.p_code
