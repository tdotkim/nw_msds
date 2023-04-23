-- SQLite
SELECT 
    VENDOR.V_NAME,
    P_CODE,
    P_DESCRIPT
FROM VENDOR
    LEFT JOIN PRODUCT ON VENDOR.V_CODE = PRODUCT.V_CODE
WHERE P_CODE IS NOT NULL
ORDER BY  V_NAME,P_CODE


