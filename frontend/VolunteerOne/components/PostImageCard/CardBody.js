import { StyleSheet, Image } from "react-native";
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
    return (
      <Block>
        <Text>
          {data["description"]}
          {"\n\n"}
        </Text>
        <Image
        style={{width: '100%', height: '50%'}}
        source={{uri:data['image']}}></Image>
      </Block>
    );
};

const styles = StyleSheet.create({
  eventName: {
    fontWeight: "bold",
    color: "#32325D",
  }
});

export default CardBody;
