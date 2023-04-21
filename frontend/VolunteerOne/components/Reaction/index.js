import { StyleSheet, View, useState } from "react-native";
import LikeButton from "./LikeButton";
import Comment from "./Comment";
import { Block, Card, Text, theme } from "galio-framework";


const Reaction = ({ likeCount, commentCount }) => {

    return (
      <View
      style={{
        borderColor: 'black',
        borderTopWidth: StyleSheet.hairlineWidth,
        borderTopWidth: .5,
        marginTop: 20,
        borderColor: "#CAD1D7"
  
      }}
    >
    <Card
        style={[styles.card]}>
      <Block flex row>

      <LikeButton likeCount={likeCount}></LikeButton>
      <Comment commentCount={commentCount}></Comment>
      </Block>
      
    </Card>
    </View>
    );
  };

  
const styles = StyleSheet.create({
    card: {
      backgroundColor: theme.COLORS.WHITE,
      width: 250,
      height: 90,
      borderRadius: 0,
      margin: 10,
      borderColor: '#fff',
      marginTop: 12,
      marginBottom: -15
    }
  });
  
  export default Reaction;
  