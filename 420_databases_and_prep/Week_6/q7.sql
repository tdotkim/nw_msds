select 
	c.cus_code as custcode,
	concat(c.cus_fname, ' ', c.cus_lname) as name,
	sum(s.sale_units*s.sale_price) as totalsales,
	t.tm_month as month
from dwcustomer as c
	inner join dwdaysalesfact as s on c.cus_code=s.cus_code
	inner join dwtime as t on s.tm_id=t.tm_id
where t.tm_month = 9
group by c.cus_code, name, month
order by totalsales desc