import { StyleSheet, View } from "react-native";
import LikeButton from "./LikeButton";
import Comment from "./Comment";
import { Block, Card, Text, theme } from "galio-framework";


const Reaction = () => {
    return (

    <Card
        style={[styles.card]}>
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
      height: 30,
      borderRadius: 0,
      margin: 10,
      borderColor: '#fff',
      marginTop: 20,
      marginBottom: -10
    }
  });
  
  export default Reaction;
  