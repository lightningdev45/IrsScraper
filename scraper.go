package main
import(
  //"fmt"
  "net/http"
  "io/ioutil"
  )

  func main(){
    //resp,err := http.Get("http://apps.irs.gov/app/picklist/list/draftTaxForms.html")
    resp,err := http.Get("http://www.irs.gov/pub/irs-dft/i8962--dft.pdf")
    if err!=nil{
      panic(err)
    }

    body, _ :=ioutil.ReadAll(resp.Body)

    // write whole the body
    err = ioutil.WriteFile("output.pdf", body, 0644)
    if err != nil {
        panic(err)
    }

    //fmt.Println(string(body))
  }
