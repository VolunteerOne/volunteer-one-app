import { StyleSheet } from "react-native";
import CardHeader from "./CardHeader";
import { Block } from "galio-framework";
import CardBody from "./CardBody";

/*
Description:
  This component returns an event card. Receives the data to insert into the card 
  as a dictionary. 

Props received:
  data - dictionary of information to insert into card
*/
const PostNoImageCard = ({ data }) => {
  return (
    <Block style={[styles.card, styles.shadowProp]}>
      <CardHeader
        author={data["author"]}
        timePosted={data["timePosted"]}
        profileImage={data["profileImage"]}
      />
      <CardBody data={data} />
    </Block>
  );
};

const styles = StyleSheet.create({
  card: {
    backgroundColor: "#FFFFFF",
    borderRadius: 8,
    padding: 15,
    marginBottom: 10,
    margin: 10
  },

  shadowProp: {
    shadowColor: "#171717",
    shadowOffset: { width: -2, height: 4 },
    shadowOpacity: 0.2,
    shadowRadius: 3,
  },
});

export default PostNoImageCard;
