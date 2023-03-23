import { StyleSheet } from "react-native";
import { Block, Text } from "galio-framework";

/*
Description:
  This component returns the body of a card. It receives the data to insert as a dictionary. 
  When receiving the data, it must have a "type" key attribute to be able to 
  dynamically generate the content of the body. 
  Example types:
    "event": organization posting an event 
    "announcment": organization posting an announcment 
Props received:
  data - dictionary of information to insert into card
*/
const CardBody = ({ data }) => {
  //retreiving type
  const type = data["type"];
  //body for an event posting
  if (type == "event") {
    return (
      <Block>
        <Text style={styles.eventName}>{data["name"]}</Text>
        <Text>
          {data["subject"]}
          {"\n\n"}
          {data["description"]}
          {"\n\n"}
          When: {data["date"]} {"\n\n"}
          Where: {data["location"]} {"\n"}
        </Text>
        <Text style={styles.eventLink}>Click to view event</Text>
      </Block>
    );
    //body for an announcment posting//
  } else if (type == "announcement") {
    return <Text>{data["announcement"]}</Text>;
    //default
  } else {
    return <Text>No data</Text>;
  }
};

const styles = StyleSheet.create({
  eventName: {
    fontWeight: "bold",
    color: "#32325D",
  },
  eventLink: {
    fontWeight: "bold",
    textDecorationLine: "underline",
    color: "#32325D",
  },
});

export default CardBody;
