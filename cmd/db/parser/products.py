import html
import json
import os
import sys
import time
import requests
from bs4 import BeautifulSoup
from googletrans import Translator

from requests.adapters import HTTPAdapter
from urllib3.util import Retry

BaseUrl = "https://www.ravenol.de/de/produkt"
# InternalUploadUrl = "http://localhost:3000/upload/"

Categories = [
    "https://www.ravenol.de/de/produkt/motorenoel/classic-motorenoel",
    "https://www.ravenol.de/de/produkt/motorenoel/2-takt-motorenoel",
    "https://www.ravenol.de/de/produkt/motorenoel/4-takt-motorenoel",
    "https://www.ravenol.de/de/produkt/motorenoel/motoroel-fuer-landwirtsch-fzg-und-baumaschinen",
    "https://www.ravenol.de/de/produkt/motorenoel/marine-motoroel",
    "https://www.ravenol.de/de/produkt/motorenoel/sonstiges-motorenoel",
    "https://www.ravenol.de/de/produkt/motorenoel/pkw-motorenoel",
    "https://www.ravenol.de/de/produkt/motorenoel/lkw-motorenoel",
    "https://www.ravenol.de/de/produkt/motorenoel/motorrad-motorenoel",

    "https://www.ravenol.de/de/produkt/getriebeoel/getriebeoel-fur-automatikgetriebe",
    "https://www.ravenol.de/de/produkt/getriebeoel/getriebeoele-fur-schaltgetriebe-und-antriebsachsen",
    "https://www.ravenol.de/de/produkt/getriebeoel/marine-getriebeol",

    "https://www.ravenol.de/de/produkt/kuehlerfrostschutz",
    "https://www.ravenol.de/de/produkt/additive",
    "https://www.ravenol.de/de/produkt/bremsflussigkeit",
    "https://www.ravenol.de/de/produkt/fahrzeugpflege-reiniger",
    "https://www.ravenol.de/de/produkt/winterchemie-zubehoer",

    "https://www.ravenol.de/de/produkt/hydraulikoel/pkw-hydraulikoel",
    "https://www.ravenol.de/de/produkt/hydraulikoel/motorrad-hydraulikoel",
    "https://www.ravenol.de/de/produkt/hydraulikoel/hydraulikole-fur-landw-fzg-und-baumaschinen",
    "https://www.ravenol.de/de/produkt/hydraulikoel/marine-hydraulikoel",
    "https://www.ravenol.de/de/produkt/hydraulikoel/sonstiges-hydraulikoel",
    
    "https://www.ravenol.de/de/produkt/industrieoel",
    "https://www.ravenol.de/de/produkt/fette",
    "https://www.ravenol.de/de/produkt/alle-werbeartikel/",
    "https://www.ravenol.de/de/produkt/zubehoer",
    "https://www.ravenol.de/de/produkt/sagekettenol"
]

Products = []
GoProducts = '''package products

import (
    uploader "main/pkg/helpers"
	"main/pkg/storage"
	"main/internal/model"
)

var AllProductsSeed = []model.Products{'''

opening = "{"
closing = "}"
translator = Translator()

def request(url, params=None, headers=None, proxies=None, timeout=10, max_retries=60, retry_delay=5):
    """
    Wrapper function for requests.get with retries, error handling, and domain resolution retries.

    Parameters:
    url (str): The URL to send the GET request to.
    params (dict, optional): URL parameters to send with the GET request.
    headers (dict, optional): Headers to send with the GET request.
    proxies (dict, optional): Proxies to use for the GET request.
    timeout (int, optional): Timeout for the GET request.
    max_retries (int, optional): Maximum number of retries for domain resolution issues.
    retry_delay (int, optional): Delay (in seconds) between retries for domain resolution issues.

    Returns:
    requests.Response: The response object, or None if the request failed.
    """
    retry_strategy = Retry(
        total=max_retries,
        status_forcelist=[429, 500, 502, 503, 504],
        allowed_methods=["HEAD", "GET", "OPTIONS"]
    )
    adapter = HTTPAdapter(max_retries=retry_strategy)
    http = requests.Session()
    http.mount("https://", adapter)
    http.mount("http://", adapter)

    for attempt in range(max_retries):
        try:
            response = http.get(url, params=params, headers=headers, proxies=proxies, timeout=timeout)
            response.raise_for_status()  # Raise an HTTPError for bad responses
            return response
        except requests.exceptions.RequestException as e:
            print(f"Attempt {attempt + 1} of {max_retries} failed: {e}")
            if isinstance(e, requests.exceptions.ConnectionError):
                # Handle specific domain resolution error
                if 'NameResolutionError' in str(e):
                    print("Domain resolution error, retrying...")
                    time.sleep(retry_delay)
                else:
                    break
            else:
                break

    print(f"Failed to complete the request after {max_retries} attempts.")
    return None

def translate(text):
    if text == "" or text == None or text == False: 
        print(f"Translation warning argument 'text' is empty", "   |   ", text)
        return ""
    for attempt in range(20):
        try:
            # Translate only if the element contains a navigable string
            translation = translator.translate(text, src="de", dest='ka').text
            break  # Break if translation was successful
        except Exception as e:
            print(f"Translation error on attempt {attempt+1}: {TimeoutError}", " \n | \n ", text)
            time.sleep((attempt + 1) * 5)  # Wait for `delay` seconds before retrying
            
    if hasattr(translation, 'text') == False:
        # print(f"Translation error object has no attribute 'text'", "   |   ", translation, "   |   ", text)
        return translation
    return translation.text

def translateHtml(html_content):
    if html_content == "" or html_content == None or html_content == False: return ""
    soup = BeautifulSoup(html_content, 'html.parser')
    text_elements = [element for element in soup.find_all(string=True) if element.parent.name not in ['script', 'style']]

    for element in text_elements:
        # Extract the text from the element
        text = element.get_text(strip=True)
        # Skip the element if it's empty
        if not text:
            continue
        # Translate the text
        translated_text = translate(text)
                
        
        # Replace the text in the element
        element.replace_with(translated_text)

    return str(soup)

headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
    "Accept-Language": "en-US,en;q=0.9",
    "Accept-Encoding": "gzip, deflate, br",
    "Connection": "keep-alive"
}

for CategoryID, Category in enumerate(Categories):
    page = 0
    while True:
        page += 1
        print(Category + "/?page=" + str(page))
        pageRawHtml = request(Category + "/?page=" + str(page), headers=headers)

        pageHtml = BeautifulSoup(pageRawHtml.text, "html.parser")
        ProductContainers = pageHtml.select("#content > div.container.pb-5 > div > div.col-lg > div > div > div > div.col-md h2.product-list__item-name a")

        if(len(ProductContainers) < 1): break

        for ProductLinks in ProductContainers:
            DetailUrl = ProductLinks.get("href")
            ProductRawHtml = request(DetailUrl)

            if hasattr(ProductRawHtml, 'text') == False:
                print("category -> ", Category, " element -> ",ProductLinks, " does not have detail url href")
                continue

            ProductHtml = BeautifulSoup(ProductRawHtml.text, "html.parser")
            print("Product " + DetailUrl + " Is Found And Fetched.")

            # @@@@@@@@@@@  DETAIL INFO  @@@@@@@@@@@


            # Title
            TitleEl = ProductHtml.select_one("#my-page > div.container > h1")
            Title = getattr(TitleEl, 'get_text', lambda: "")() or ""





            # RIGHT SIDE DESCRIPTION
            RightSide = ProductHtml.select_one(".product__right") 
            Desc = translateHtml(str(RightSide))






            # TABLE
            imagesNames = ProductHtml.select("#content > div > div.product__left.pe-0.pe-md-4 > div > div > div >  div.product__option-item")
            images = f'''[]model.Files{opening}'''

            for image in imagesNames:
                imageTitle = image.get("data-title")
                img = ProductHtml.select_one("#content [data-title='" + imageTitle + "'] img")

                # https://www.ravenol.de/storage/app/uploads/public/d0f/2b5/96a/thumb__600_0_0_0_crop.jpg
                imgSrc = img.get("src").replace("__100_", "__600_")
                print(imgSrc)
                response = request(imgSrc)

                # Check if the request was successful (status code 200)
                if response.status_code != 200:
                    print("Product " + Title + " Failed To Download " + imgSrc, response.status_code)
                else:
                    basename, extension = os.path.splitext(imgSrc)
                    fileName = basename.replace("https://www.ravenol.de/storage/app/uploads/public/", "").replace("/", "-") + extension

                    with open("./public/uploads/products/images/" + fileName, "wb") as file:
                        file.write(response.content)
                    
                    with open("./public/uploads/products/images/" + fileName, 'rb') as file:
                        images += f'''\n            {opening}
                Name:       `{fileName}`,
                Original:   `{imageTitle}`,
                Location:   "local",
                Path:       "/uploads/products/images/{fileName}",
                Size:       {file.__sizeof__()},
                Base64:     "",
                Compressed: false,
                TypeID:     uploader.GetDbTypeIdByExtension(`{extension.replace(".", "")}`),
            {closing}, '''
                        # files = {'file': file}
                        # Send the file using a POST request
                        # ImgUploadResponse = requests.post(InternalUploadUrl, files=files)
                        # ThumbnailID = int(ImgUploadResponse.json()["ID"])

            images += f'''\n        {closing},'''





            # TABLE
            tableSelector = "#content > div > div.product__right > div.product__desc > div.table-responsive > table tbody tr"
            tableRows = ProductHtml.select(tableSelector)
            table = f'''[]model.Product_properties{opening}'''
            tableData = []

            for rowID, tableRow in enumerate(tableRows):
                characteristics = ProductHtml.select_one(tableSelector + ":nth-child(" + str(rowID + 1) + ") td:nth-child(1)")
                unit = ProductHtml.select_one(tableSelector + ":nth-child(" + str(rowID + 1) + ") td:nth-child(2)")
                data = ProductHtml.select_one(tableSelector + ":nth-child(" + str(rowID + 1) + ") td:nth-child(3)")
                examination = ProductHtml.select_one(tableSelector + ":nth-child(" + str(rowID + 1) + ") td:nth-child(4)")
                
                row = {
                    "characteristics": getattr(characteristics, 'get_text', lambda: "")() or "",
                    "unit": getattr(unit, 'get_text', lambda: "")() or "",
                    "data": getattr(data, 'get_text', lambda: "")() or "",
                    "examination": getattr(examination, 'get_text', lambda: "")() or "",
                }

                table += f'''
            {opening}
                Characteristics: `{translate(row["characteristics"])}`,
                Unit: `{translate(row["unit"])}`,
                Data: `{translate(row["data"])}`,
                Examination: `{translate(row["examination"])}`,
            {closing},'''
        
                tableData.append(row)

            table += f'''\n        {closing},'''




            # PDF
            pdfURL = ProductHtml.select_one("#content > div > a").get("href")
            # print("https://www.ravenol.de" + pdfURL)
            pdfResponse = request("https://www.ravenol.de" + pdfURL)
            pdf = f'''model.Files{opening}'''

            if pdfResponse.status_code != 200:
                print("Product " + Title + " PDF -> Failed To Download " + "https://www.ravenol.de" + pdfURL, pdfResponse.status_code)
            else:
                basename, extension = os.path.splitext(pdfURL)
                parts = basename.split("/")
                fileName = parts[-1] + extension

                with open("./public/uploads/products/pdfs/" + fileName, "wb") as file:
                    file.write(pdfResponse.content)
                
                with open("./public/uploads/products/pdfs/" + fileName, 'rb') as file:
                    pdf += f'''
            Name:       `{fileName}`,
            Original:   `{fileName}`,
            Location:   "local",
            Path:       "/uploads/products/pdfs/{fileName}",
            Size:       {file.__sizeof__()},
            Base64:     "",
            Compressed: false,
            TypeID:     uploader.GetDbTypeIdByExtension(`{extension.replace(".", "")}`),'''
            pdf += f'''\n        {closing},'''

            print("Prodcut " + Title + " Is Parsed.")


            GoProducts += f'''\n{opening}
        Name: `{Title}`,
        Slug: `{Title.replace(" ", "_")}`,
        Description: `{Desc}`,
        DescriptionHtml: `{Desc}`,
        CategoryID: {CategoryID + 1},
        TechnicalSheet: {pdf}
        Pics: {images}
        Properties: {table}
    {closing},
'''

            time.sleep(3)

            # break
GoProducts += '''}

func Populate() {
    for _, row := range AllProductsSeed { storage.DB.Create(&row) }
}
'''

# cleaning
GoProducts.replace("™", "")


with open("./cmd/db/seed/products/products.seed.go", "w", encoding='utf-8') as file:
    file.write(GoProducts)

# with open("./cmd/parser/result.json", "w") as json_file:
#     json.dump(Products, json_file, indent=4)
