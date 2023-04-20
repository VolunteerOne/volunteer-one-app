import { StyleSheet, Dimensions } from "react-native";
import { Block, Text, Button, theme } from "galio-framework";
import { useNavigation } from "@react-navigation/native";
/*
Description:
  This component returns the body of a card. It receives the data to insert as a dictionary. 
  When receiving the data, it must have a "type" key attribute to be able to 
  dynamically generate the content of the body. 
  Example types:
    "event":        organization posting an event. For event posts there is a button
                    that a user can click to view the event details
    "announcment": organization posting an announcment 
Props received:
  data - dictionary of information to insert into card
*/
const { height, width } = Dimensions.get("window");
const CardBody = ({ data }) => {
  //retreiving type
  const type = data["type"];
  const navigation = useNavigation();

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
        <Button
          style={styles.button}
          shadowless
          onPress={() =>
            navigation.navigate("ViewEvent", { eventID: data["id"] })
          }
        >
          <Text style={styles.eventLink}>Click to view event</Text>
        </Button>
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
  button: {
    backgroundColor: theme.COLORS.TRANSPARENT,
    width: width * 0.35,
    borderRadius: 0,
    borderWidth: 0,
    height: 24,
    elevation: 0,
  },
});

export default CardBody;
