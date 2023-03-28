import { StyleSheet, View } from "react-native";
import LikeButton from "./LikeButton";
import Comment from "./Comment";
import { Block, Card, Text, theme } from "galio-framework";

import argonTheme from "../../constants";

const Reaction = () => {
    return (
    <Card
        style={[styles.card, styles.shadowProp]}>
      <Block flex row>
      <LikeButton></LikeButton>
      <Comment></Comment>
      </Block>
    </Card>
    );
  };

  
const styles = StyleSheet.create({
    card: {
      backgroundColor: theme.COLORS.WHITE,
      width: 400,
      height: 40,
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
  
  export default Reaction;
  