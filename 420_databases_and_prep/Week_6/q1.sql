select 
	r.reg_name, 
	s.cus_code, 
	c.cus_lname,
	c.cus_fname,
	sum(s.sale_units*s.sale_price) as sales
from dwcustomer as c
	inner join dwdaysalesfact as s on c.cus_code=s.cus_code
	inner join dwregion as r on c.reg_id=r.reg_id
group by r.reg_name, s.cus_code,c.cus_lname, c.cus_fname
order by r.reg_name