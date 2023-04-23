-- SQLite
SELECT 
    CUSTOMER.CUS_CODE,
    CUSTOMER.CUS_LNAME,
    CUSTOMER.CUS_BALANCE
FROM CUSTOMER
    OUTER JOIN INVOICE ON INVOICE.CUS_CODE = CUSTOMER.CUS_CODE
GROUP BY INVOICE.CUS_CODE
ORDER BY CUSTOMER.CUS_CODE





