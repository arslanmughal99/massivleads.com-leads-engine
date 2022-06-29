<div align="center">
<h1>
Lead Extracting engine for Massivleads.com
</h1>
</div>

This is the original leads extracting engine used to extract leads to generate facebook custom audience and share with massivleads premium clients to get targeted audience .

### Usage:

1. Clone repository
  
2. update .env file in project root
  
3. Install packages and run program
  
  ```bash
  go run .
  ```
  
4. put a post request to localhost:3002/scraper with following data as json or formdata .
  
  ```json
  {
  "id": "ANY_UNIQUE_ID_UUID4",
  "domains": "DOMAINS_TO_EXTRACT_EMAILS_FOR",
  "keyword": "KEYWORD_THAT_NEED_TO_INCLUDE_IN_SEARCH",
  "country": "TARGET_COUNTRY",
  "jobTitle": "JOB_TITLES_TO_EXTRACT_TARGET_LEADS_FOR"
  }
  ```
  

After posting data engine will extract leads and post it to your webhook URL provided in .env file .