select 
	coalesce(r.reg_name,'grand total') as reg_name,
	coalesce(cast(s.cus_code as varchar),'region_subtotal') as cus_code,
	coalesce(concat(c.cus_fname, ' ', c.cus_lname), 'total') as cus_fullname,
	sum(s.sale_units*s.sale_price) as sales
from dwcustomer as c
	inner join dwdaysalesfact as s on c.cus_code=s.cus_code
	inner join dwregion as r on c.reg_id=r.reg_id
group by rollup (r.reg_name, s.cus_code, c.cus_fname,c.cus_lname)
order by r.reg_name