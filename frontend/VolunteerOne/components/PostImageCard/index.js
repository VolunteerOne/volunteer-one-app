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
      </Block>
      <Reaction likeCount={data["likeCount"]}
      commentCount={data["commentCount"]} />
    </Block>
  );
};

const styles = StyleSheet.create({
  card: {
    backgroundColor: "#FFFFFF",
    width: "100%",
    minWidth: "100%",
    borderRadius: 10,
    padding: 15,
    margin: 10,
  },
  shadowProp: {
    shadowColor: "#171717",
    shadowOffset: { width: -2, height: 4 },
    shadowOpacity: 0.2,
    shadowRadius: 3,
  },
});

export default PostImageCard;
