import { StyleSheet, View } from "react-native";
import LikeButton from "./LikeButton";
import Comment from "./Comment";
import { Block, Card, Text, theme } from "galio-framework";


const Reaction = () => {
    return (
      <View
      style={{
        borderColor: 'black',
        borderTopWidth: StyleSheet.hairlineWidth,
        borderTopWidth: .5,
        marginTop: 10,
        borderColor: "#32325D"
      }}
    >
    <Card
        style={[styles.card]}>


      <Block flex row>

      <LikeButton></LikeButton>
      <Comment></Comment>
      </Block>
      
    </Card>
    </View>
    );
  };

  
const styles = StyleSheet.create({
    card: {
      backgroundColor: theme.COLORS.WHITE,
      width: 250,
      height: 30,
      borderRadius: 0,
      margin: 10,
      borderColor: '#fff',
      marginTop: 15,
      marginBottom: -15
    }
  });
  
  export default Reaction;
  