select 
	coalesce(r.reg_name,'...grand total') as reg_name,
	coalesce(cast(s.cus_code as varchar),'...region_subtotal') as cus_code,
	coalesce(cus_fullname, '...cuscode subtotal') as cus_fullname,
	sum(s.sale_units*s.sale_price) as sales
from 
	(select
	 	dwcustomer.cus_code,
		concat(dwcustomer.cus_fname, ' ', dwcustomer.cus_lname) as cus_fullname,
	 	dwcustomer.reg_id
	 from
	 	dwcustomer
	) as c
	inner join dwdaysalesfact as s on c.cus_code=s.cus_code
	inner join dwregion as r on c.reg_id=r.reg_id
group by cube (r.reg_name, s.cus_code, c.cus_fullname)
order by r.reg_name, s.cus_code,c.cus_fullname asc;