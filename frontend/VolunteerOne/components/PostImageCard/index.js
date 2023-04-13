import { StyleSheet } from "react-native";
import CardHeader from "./CardHeader";
import { Block } from "galio-framework";
import CardBody from "./CardBody";
import Reaction from "../Reaction";
import React from "react";
/*
Description:
  This component returns an event card. Receives the data to insert into the card 
  as a dictionary. 

Props received:
  data - dictionary of information to insert into card
*/
const PostImageCard = ({ data }) => {
  return (
    <Block style={[styles.card, styles.shadowProp]}>
      <CardHeader
        author={data["author"]}
        timePosted={data["timePosted"]}
        profileImage={data["profileImage"]}
      />
      <CardBody data={data} />
      <Block>
      <Reaction></Reaction>
      </Block>
    </Block>

  );
};

const styles = StyleSheet.create({
  card: {
    backgroundColor: "#FFFFFF",
    borderRadius: 8,
    padding: 15,
    marginBottom: 10,
    margin: 10,
    width: 'auto'
  },

  shadowProp: {
    shadowColor: "#171717",
    shadowOffset: { width: -2, height: 4 },
    shadowOpacity: 0.2,
    shadowRadius: 3,
  },
});

export default PostImageCard;
