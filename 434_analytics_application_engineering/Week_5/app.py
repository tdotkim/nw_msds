import pandas as pd
import sys
import numpy as np
import os
from google.cloud import bigquery
from flask import Flask, render_template, request, escape

app = Flask(__name__)

#ONLY FOR LOCAL USE
#CREDS = "C:\\Users\\TK\Desktop\\msds434-module5-af0b121e8d76.json"
#client = bigquery.Client.from_service_account_json(json_credentials_path=CREDS,project='msds434-module5')
client = bigquery.Client(project='msds434-module5')

dset = 'module5'

@app.route('/')
def index():
    return render_template('index.html')

@app.route('/set_base_model')
def form():
    return render_template('form.html')

@app.route('/data', methods = ['POST', 'GET'])
def data():
    if request.method == 'POST':
        form_data = request.form
        start = str(form_data['Start'])
        end = str(form_data['End'])
        pstart = str(form_data['Prediction Start'])
        pend = str(form_data['Prediction End'])
        name = str(form_data['Model Name'])
        query = """
                CREATE OR REPLACE MODEL {0}.{1}
                OPTIONS(model_type='logistic_reg',
                        input_label_cols=['isBuyer'])
                AS
                SELECT
                    IF(totals.transactions IS NULL, 0, 1) as isBuyer,
                    IFNULL(totals.pageviews,0) AS pageviews,
                    IFNULL(totals.timeOnSite,0) AS timeOnSite,
                    IFNULL(totals.newVisits,0) AS isNewVisit,
                    IF(device.deviceCategory = 'mobile', 1, 0) as isMobile,
                    IF(device.deviceCategory = 'desktop', 1, 0) as isDesktop,
                    IF(trafficSource.medium in ('affiliate','cpc','cpm'),1,0) AS isPaidTraffic
                FROM
                    `bigquery-public-data.google_analytics_sample.ga_sessions_*`
                WHERE
                    _TABLE_SUFFIX BETWEEN @start AND @end
            """.format(dset, name)


        job_config = bigquery.QueryJobConfig(
            query_parameters=[
                bigquery.ScalarQueryParameter("start", "STRING", start),
                bigquery.ScalarQueryParameter("end", "STRING", end),
            ]
        )

        query_job = client.query(query, job_config=job_config)  

        try: 
            for row in query_job:
                print(row)
        except:
            print("ignoring this error: {}".format(sys.exc_info())) #  the model was successfully created

        # do the eval
        query = """
        SELECT *
        FROM ML.EVALUATE(MODEL {0}.{1},
        (
        SELECT
        IF(totals.transactions IS NULL, 0, 1) as isBuyer,
        IFNULL(totals.pageviews,0) AS pageviews,
        IFNULL(totals.timeOnSite,0) AS timeOnSite,
        IFNULL(totals.newVisits,0) AS isNewVisit,
        IF(device.deviceCategory = 'mobile', 1, 0) as isMobile,
        IF(device.deviceCategory = 'desktop', 1, 0) as isDesktop,
        IF(trafficSource.medium in ('affiliate','cpc','cpm'),1,0) AS isPaidTraffic
        FROM
        `bigquery-public-data.google_analytics_sample.ga_sessions_*`
        WHERE
        _TABLE_SUFFIX BETWEEN @start AND @end
        ),
        STRUCT(0.5 AS threshold)
        )
        """.format(dset, name)


        job_config = bigquery.QueryJobConfig(
            query_parameters=[
                bigquery.ScalarQueryParameter("start", "STRING", start),
                bigquery.ScalarQueryParameter("end", "STRING", end),
            ]
        )

        df_eval = client.query(query, job_config=job_config).to_dataframe()

        # do the predict
        query = """
        SELECT *
        FROM ML.PREDICT(MODEL {0}.{1},
        (
        SELECT
        IF(totals.transactions IS NULL, 0, 1) as isBuyer,
        IFNULL(totals.pageviews,0) AS pageviews,
        IFNULL(totals.timeOnSite,0) AS timeOnSite,
        IFNULL(totals.newVisits,0) AS isNewVisit,
        IF(device.deviceCategory = 'mobile', 1, 0) as isMobile,
        IF(device.deviceCategory = 'desktop', 1, 0) as isDesktop,
        IF(trafficSource.medium in ('affiliate','cpc','cpm'),1,0) AS isPaidTraffic
        FROM
        `bigquery-public-data.google_analytics_sample.ga_sessions_*`
        WHERE
        _TABLE_SUFFIX BETWEEN @start AND @end
        ),
        STRUCT(0.5 AS threshold)
        )
        """.format(dset, name)


        job_config = bigquery.QueryJobConfig(
            query_parameters=[
                bigquery.ScalarQueryParameter("start", "STRING", start),
                bigquery.ScalarQueryParameter("end", "STRING", end),
            ]
        )

        df = client.query(query, job_config=job_config).to_dataframe()
        df2 = pd.DataFrame(df.explode('predicted_isBuyer_probs')['predicted_isBuyer_probs'])
        df2 = df2['predicted_isBuyer_probs'].apply(pd.Series)
        pivoted = df2.pivot_table(index=df2.index,columns='label',values='prob',aggfunc=np.mean)
        pivoted.rename(columns={ pivoted.columns[0]: "Prob_Not_Buyer" ,pivoted.columns[1]: "Prob_IS_Buyer" }, inplace = True)
        merged = df.drop('predicted_isBuyer_probs',axis=1)
        merged = merged.join(pivoted)
        merged
        return render_template('data.html',tables=[df_eval.to_html(classes='data'),merged.to_html(max_rows=20,classes='data')], titles=['eval scores','predictions'])
    
 
 
if __name__ == "__main__":
    app.run(debug=True, host="0.0.0.0", port=int(os.environ.get("PORT", 8080)))
