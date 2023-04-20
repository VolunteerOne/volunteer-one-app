// This page is temporarily being used to test other pages (notifications, friends, etc.) - matt
import React from 'react';
import { StyleSheet, Dimensions, ScrollView } from 'react-native';
import { Block, theme } from 'galio-framework';
import ProfileItem from '../../components/ProfileItem';
import profile from '../../constants/ProfileTab/profile';

// ================================= View Friends Page ================================= //

const { width } = Dimensions.get('screen');

class ViewFriendsPage extends React.Component {
  renderFriends = () => {
    // determine the user 
    /* 1. Get the props */
    const {navigation,route}=this.props
    /* 2. Get the param */
    theUser = route.params.theUser;   // determine which profile you are viewing
    console.log("Viewing friends of: ", theUser)


    return (
      <ScrollView
        showsVerticalScrollIndicator={false}
        contentContainerStyle={styles.notifications}>
        
        {/* map all notifications using loop */}
        <Block flex>
          {profile[theUser].friends.friendsList.map((profile,i) => (
             <ProfileItem key={i} item={profile} horizontal />
          ))}
        </Block>
    
      </ScrollView>
    )
  }

  render() {
    return (
      <Block flex center style={styles.home}>
        {this.renderFriends()}
      </Block>
    );
  }
}

const styles = StyleSheet.create({
  home: {
    width: width,    
  },
  notifications: {
    width: width - theme.SIZES.BASE * 2,
    paddingVertical: theme.SIZES.BASE,
  },
});

export default ViewFriendsPage;
