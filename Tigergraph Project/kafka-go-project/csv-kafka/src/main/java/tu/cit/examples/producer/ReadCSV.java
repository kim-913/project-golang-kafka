package tu.cit.examples.producer;


import com.opencsv.CSVReader;
import com.opencsv.bean.CsvToBean;
import com.opencsv.bean.CsvToBeanBuilder;

import java.io.FileReader;
import java.util.List;

public class ReadCSV {
    public String csvFileName = "data/data.csv";
    public List userLst;
    public List ReadCSVFile()  {

        try {
            CSVReader csvReader = new CSVReader(new FileReader(csvFileName));

            CsvToBean csvToBean = new CsvToBeanBuilder(csvReader)
                    .withType(userModel.class)
                    .withIgnoreLeadingWhiteSpace(true).build();

            userLst = csvToBean.parse();
            csvReader.close();
        }catch(Exception FileNotFoundException){
            //e.printStackTrace();
            System.out.println("File is not available...");
        }

        return userLst;
    }
}